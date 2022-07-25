package models

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/mmuiro/riichi-base/src/models/constants/groups"
	"github.com/mmuiro/riichi-base/src/models/constants/suits"
	"github.com/mmuiro/riichi-base/src/models/constants/waits"
	"github.com/mmuiro/riichi-base/src/utils"
)

type Hand struct {
	ClosedTiles        []Tile
	Tenpai             bool
	Waits              map[int][]Partition
	Melds              []Mentsu
	effectiveTileCount int
}

func (h *Hand) Tiles() []Tile {
	tiles := make([]Tile, 0)
	for _, tile := range h.ClosedTiles {
		tiles = append(tiles, tile)
	}
	for _, mentsu := range h.Melds {
		for _, tile := range mentsu.Tiles {
			tiles = append(tiles, tile)
		}
	}
	return tiles
}

func (h *Hand) String() string {
	tileStrings := make([]string, len(h.Tiles()))
	for i, tile := range h.Tiles() {
		tileStrings[i] = tile.String()
	}
	return strings.Join(tileStrings, " ")
}

type InvalidHandError struct{}

func (e *InvalidHandError) Error() string {
	return "A hand must be 13 tiles or complete (14 tiles), without counting extras from kans."
}

type InvalidHonorError struct{}

func (e *InvalidHonorError) Error() string {
	return "Honor tiles must be numbered from 1 to 7."
}

type RedDoraError struct{}

func (e *RedDoraError) Error() string {
	return "There can only be one red five per suit."
}

type MissingTileError struct {
	t *Tile
}

func (e *MissingTileError) Error() string {
	return fmt.Sprintf("The tile %s was not found.", e.t.String())
}

func CreateHand(tiles []Tile, melds []Mentsu) *Hand {
	sort.Slice(tiles, func(i, j int) bool {
		return TileToID(&tiles[i]) < TileToID(&tiles[j])
	})
	h := &Hand{ClosedTiles: tiles, Melds: melds}
	return h
}

/* Takes a string and coverts it into a hand. */
func StringToHand(s string) (*Hand, *Tile, error) {
	tileCount, melds := 0, make([]Mentsu, 0)
	var closedTiles []Tile
	var err error
	kans, redCounts := 0, make(map[suits.Suit]int)
	r := regexp.MustCompile(`\s+`)
	parts := r.Split(s, -1)
	var lastTile Tile
	for i, part := range parts {
		if i == 0 {
			// closed portion of the hand
			closedTiles, _, err = parsePartToTiles(part, redCounts, false)
			if err != nil {
				return nil, nil, err
			}
			tileCount += len(closedTiles)
			lastTile = closedTiles[len(closedTiles)-1]
		} else {
			// melds (pon, chii). kans can be specified as closed; otherwise, they are open.
			var meldTiles []Tile
			var open bool
			var err error
			meldTiles, open, err = parsePartToTiles(part, redCounts, true)
			tileCount += len(meldTiles)
			if err != nil {
				return nil, nil, err
			}
			var newMeld *Mentsu
			newMeld, err = CreateMentsu(meldTiles, open)
			if err != nil {
				return nil, nil, err
			}
			if newMeld.Kind == groups.Kantsu {
				kans++
			}
			melds = append(melds, *newMeld)
		}
	}
	// still need to deal with kans...
	effectiveTileCount := tileCount - kans
	if effectiveTileCount != 13 && effectiveTileCount != 14 {
		return nil, nil, &InvalidHandError{}
	}
	hand := CreateHand(closedTiles, melds)
	if effectiveTileCount == 14 {
		return hand, &lastTile, nil
	}
	return hand, nil, nil
}

func parsePartToTiles(part string, redCounts map[suits.Suit]int, defaultOpen bool) ([]Tile, bool, error) {
	tiles := make([]Tile, 0)
	values := make([]int, 0)
	open := defaultOpen
	for _, c := range part {
		suit, ok := suits.CharToSuit[c]
		if !ok {
			if c == 'z' {
				// honor tile
				for _, v := range values {
					if v < 1 || v > 7 {
						return nil, open, &InvalidHonorError{}
					}
					suit = suits.Suit(v + 2)
					tiles = append(tiles, *CreateTile(suit, 0, false))
				}
				values = make([]int, 0)
			} else if c == 'c' {
				// indicating closed kan (only used on parts of the second portion of string)
				open = false
			} else {
				values = append(values, int(c-'0'))
			}
		} else {
			for _, v := range values {
				var newTile *Tile
				if v == 0 {
					// red five
					newTile = CreateTile(suit, 5, true)
					redCounts[suit]++
					if redCounts[suit] > 1 {
						return nil, open, &RedDoraError{}
					}
				} else {
					newTile = CreateTile(suit, v, false)
				}
				tiles = append(tiles, *newTile)
			}
			values = make([]int, 0)
		}
	}
	return tiles, open, nil
}

func getTileIndex(tiles []Tile, suit suits.Suit, value int) int {
	var l, h int = 0, len(tiles) - 1
	var m int
	for h >= l {
		m = (h + l) / 2
		if tiles[m].Value == value && tiles[m].Suit == suit {
			return m
		} else if SuitAndValueToID(suit, value) > TileToID(&tiles[m]) {
			l = m + 1
		} else {
			h = m - 1
		}
	}
	return -1
}

func TilesToString(tiles []Tile) string {
	var ret string
	prevSuit, wasHonor := tiles[0].Suit, tiles[0].IsHonor()
	for _, tile := range tiles {
		if tile.Suit != prevSuit && (!tile.IsHonor() || !wasHonor) {
			if suit, ok := suits.SuitToString[prevSuit]; ok {
				ret += suit
			} else {
				ret += "z"
			}
			prevSuit = tile.Suit
			wasHonor = tile.IsHonor()
		}
		if tile.Red {
			ret += "5"
		} else if !tile.IsHonor() {
			ret += strconv.Itoa(tile.Value)
		} else {
			ret += strconv.Itoa(int(tile.Suit) - 2)
		}
	}
	if wasHonor {
		ret += "z"
	} else {
		ret += suits.SuitToString[prevSuit]
	}
	return ret
}

/* Checks if the given hand, on the last tile t, is a winning hand (和了).
Assumes the hand's waits are already set from calling CheckTenpai. */
func CheckAgari(h *Hand, t *Tile, tsumo bool) (bool, []Partition) {
	if !h.Tenpai {
		return false, nil
	}
	tileID := TileToID(t)
	agariPartitions := make([]Partition, 0)
	if h.Waits[tileID] != nil {
		for _, p := range h.Waits[TileToID(t)] {
			if p.Wait == waits.KokushiSingle {
				newMentsu, _ := CreateMentsu([]Tile{*t}, false)
				p.Mentsu = append(p.Mentsu, *newMentsu)
			} else {
				var condition func(m *Mentsu) bool
				switch p.Wait {
				case waits.Ryanmen:
					condition = func(m *Mentsu) bool {
						return m.Kind == groups.Ryanmen
					}
				case waits.Kanchan:
					condition = func(m *Mentsu) bool {
						return m.Kind == groups.Kanchan
					}
				case waits.Penchan:
					condition = func(m *Mentsu) bool {
						return m.Kind == groups.Penchan
					}
				case waits.Shanpon:
					condition = func(m *Mentsu) bool {
						return m.Kind == groups.Toitsu && m.Tiles[0].Equals(t)
					}
				case waits.Tanki:
					condition = func(m *Mentsu) bool {
						return m.Kind == groups.Single && m.Tiles[0].Equals(t)
					}
				case waits.KokushiThirteen:
					condition = func(m *Mentsu) bool {
						return m.Kind == groups.Single && m.Tiles[0].Equals(t)
					}
				}
				for i := range p.Mentsu {
					if CheckJunseiChuuren(&p) {
						p.Wait = waits.JunseiChuuren
					}
					if condition(&p.Mentsu[i]) {
						p.Mentsu[i].addTile(t)
						if !tsumo {
							p.Mentsu[i].Open = true
						}
						break
					}
				}
				AssignMentsuCounts(&p)
			}
			// needed or not needed based on value/reference passing...
			agariPartitions = append(agariPartitions, p)
		}
		return true, agariPartitions
	}
	return false, nil
}

func CheckComplete(h *Hand) (bool, []Partition) {
	if len(h.Tiles()) < 14 {
		return false, nil
	}
	// checking arbitrary hands (i.e. from calculator)
	completePartitions := make([]Partition, 0)
	complete := false
	for _, partition := range CalculateAllPartitions(h) {
		if CheckStandard(&partition) || CheckChiiToitsu(&partition) || CheckKokushi(&partition) {
			completePartitions = append(completePartitions, partition)
			complete = true
		}
	}
	return complete, completePartitions
}

func CheckTenpai(h *Hand) (bool, []Partition, [][]int) {
	tenpaiPartitions, tenpaiWaits := make([]Partition, 0), make([][]int, 0)
	tenpai := false
	checks := []func(p *Partition) (bool, []int){CheckRyanmen, CheckKanchan, CheckPenchan, CheckShanpon, CheckTanki, CheckKokushiSingle, CheckKokushiThirteen}
	for _, partition := range CalculateAllPartitions(h) {
		curWaits := make([]int, 0)
		for i, check := range checks {
			if passed, waitTileIDs := check(&partition); passed {
				tenpai = true
				partition.Wait = waits.WaitKind(i)
				curWaits = append(curWaits, waitTileIDs...)
				tenpaiPartitions = append(tenpaiPartitions, partition)
				tenpaiWaits = append(tenpaiWaits, curWaits)
				break
			}
		}
	}
	return tenpai, tenpaiPartitions, tenpaiWaits
}

func AssignWaitMap(h *Hand, partitions []Partition, waitsList [][]int) error {
	h.Waits = make(map[int][]Partition)
	for i, p := range partitions {
		for _, id := range waitsList[i] {
			h.Waits[id] = append(h.Waits[id], p)
		}
	}
	return nil
}

// Remove a tile from the closed portion of a hand.
func (h *Hand) RemoveTile(t *Tile) error {
	if i := getTileIndex(h.ClosedTiles, t.Suit, t.Value); i > -1 {
		h.ClosedTiles = utils.RemoveIndex(h.ClosedTiles, i)
		// check if the hand is still valid
		return nil
	}
	return &MissingTileError{}
}

// Add a tile to the closed portion of a hand.
func (h *Hand) AddTile(t *Tile) error {
	h.ClosedTiles = append(h.ClosedTiles, *t)
	sort.Slice(h.ClosedTiles, func(i, j int) bool {
		return TileToID(&h.ClosedTiles[i]) < TileToID(&h.ClosedTiles[j])
	})
	// check if the hand is still valid
	return nil
}
