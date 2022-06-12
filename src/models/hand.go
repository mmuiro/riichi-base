package models

import (
	"regexp"
	"riichi-calculator/src/models/constants/groups"
	"riichi-calculator/src/models/constants/suits"
	"riichi-calculator/src/models/constants/waits"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	Tiles  []Tile
	Tenpai bool
	waits  map[int][]Partition
	Melds  []Mentsu
}

func (h Hand) String() string {
	tileStrings := make([]string, len(h.Tiles))
	for i, tile := range h.Tiles {
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

func CreateHand(tiles []Tile, melds []Mentsu) *Hand {
	sort.Slice(tiles, func(i, j int) bool {
		return TileToID(&tiles[i]) < TileToID(&tiles[j])
	})
	h := &Hand{Tiles: tiles, Melds: melds}
	return h
}

/* Takes a string and coverts it into a hand. */
func StringToHand(s string) (*Hand, error) {
	tiles, melds := make([]Tile, 0), make([]Mentsu, 0)
	kans, redCounts := 0, make(map[suits.Suit]int)
	r := regexp.MustCompile(`\s+`)
	parts := r.Split(s, -1)
	for i, part := range parts {
		if i == 0 {
			// closed portion of the hand
			closedTiles, _, err := parsePartToTiles(part, redCounts, false)
			if err != nil {
				return nil, err
			}
			tiles = append(tiles, closedTiles...)
		} else {
			// melds (pon, chii). kans can be specified as closed; otherwise, they are open.
			var meldTiles []Tile
			var open bool
			var err error
			meldTiles, open, err = parsePartToTiles(part, redCounts, true)
			tiles = append(tiles, meldTiles...)
			if err != nil {
				return nil, err
			}
			var newMeld *Mentsu
			newMeld, err = CreateMentsu(meldTiles, open)
			if err != nil {
				return nil, err
			}
			melds = append(melds, *newMeld)
		}
	}
	effectiveTileCount := len(tiles) - kans
	if effectiveTileCount != 13 && effectiveTileCount != 14 {
		return nil, &InvalidHandError{}
	}
	return CreateHand(tiles, melds), nil
}

func parsePartToTiles(part string, redCounts map[suits.Suit]int, defaultOpen bool) ([]Tile, bool, error) {
	tiles := make([]Tile, 0)
	values := make([]int, 0)
	open := defaultOpen
	for _, c := range part {
		suit, ok := suits.CharToSuit[c]
		if !ok && c == 'z' {
			// honor tile
			for _, v := range values {
				if v < 1 || v > 7 {
					return nil, open, &InvalidHonorError{}
				}
				suit = suits.Suit(v + 2)
				tiles = append(tiles, *CreateTile(suit, 0, false))
			}
			values = make([]int, 0)
		} else if !ok {
			values = append(values, int(c-'0'))
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

func CheckAgari(h *Hand, t *Tile) (bool, []Partition) {
	if !h.Tenpai {
		return false, nil
	}
	tileID := TileToID(t)
	agariPartitions := make([]Partition, 0)
	if h.waits[tileID] != nil {
		for _, partition := range h.waits[TileToID(t)] {
			if partition.Wait == waits.KokushiSingle {
				newMentsu, _ := CreateMentsu([]Tile{*t}, false)
				partition.Mentsu = append(partition.Mentsu, *newMentsu)
			} else {
				var condition func(m *Mentsu) bool
				switch partition.Wait {
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
				for _, m := range partition.Mentsu {
					if condition(&m) {
						m.addTile(t)
						break
					}
				}
			}
			// needed or not needed based on value/reference passing...
			agariPartitions = append(agariPartitions, partition)
		}
		return true, agariPartitions
	}
	return false, nil
}

func CheckComplete(h *Hand) (bool, []Partition) {
	if len(h.Tiles) < 14 {
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

func CheckTenpai(h *Hand) (bool, map[int][]Partition) {
	waitMap := make(map[int][]Partition)
	tenpai := false
	checks := []func(p *Partition) (bool, []int){CheckRyanmen, CheckKanchan, CheckPenchan, CheckShanpon, CheckTanki, CheckKokushiSingle, CheckKokushiThirteen}
	for _, partition := range CalculateAllPartitions(h) {
		for i, check := range checks {
			if passed, waitTileIDs := check(&partition); passed {
				tenpai = true
				partition.Wait = waits.WaitKind(i)
				for _, id := range waitTileIDs {
					waitMap[id] = append(waitMap[id], partition)
				}
				break
			}
		}
	}
	return tenpai, waitMap
}
