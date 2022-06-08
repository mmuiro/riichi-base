package models

import (
	"riichi-calculator/src/utils"
	"sort"
	"strconv"
	"strings"
)

type TileCounts [34]int

type Partition struct {
	Mentsu []Mentsu
}

func (p Partition) String() string {
	mentsuStrings := make([]string, len(p.Mentsu))
	for i, mentsu := range p.Mentsu {
		mentsuStrings[i] = mentsu.String()
	}
	return strings.Join(mentsuStrings, "  ")
}

func (p Partition) TotalTiles() int {
	sum := 0
	for _, p := range p.Mentsu {
		sum += len(p.tiles)
	}
	return sum
}

type Hand struct {
	tiles  []Tile
	tenpai bool
	waits  []Tile
	melds  []Mentsu
}

func (h Hand) String() string {
	tileStrings := make([]string, len(h.tiles))
	for i, tile := range h.tiles {
		tileStrings[i] = tile.String()
	}
	return strings.Join(tileStrings, " ")
}

func (h Hand) Tiles() []Tile {
	return h.tiles
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

var charToSuit = map[rune]Suit{
	'm': Man,
	's': Sou,
	'p': Pin,
}

func CreateHand(tiles []Tile, melds []Mentsu) *Hand {
	sort.Slice(tiles, func(i, j int) bool {
		return TileToID(&tiles[i]) < TileToID(&tiles[j])
	})
	h := &Hand{tiles: tiles, melds: melds}
	return h
}

/* Takes a string and coverts it into a hand. */
func StringToHand(s string) (*Hand, error) {
	tiles, melds := make([]Tile, 0), make([]Mentsu, 0)
	kans, redCounts := 0, make(map[Suit]int)
	parts := strings.Split(s, " ")
	for i, part := range parts {
		if i == 0 {
			// closed portion of the hand
			values := make([]int, 0)
			for _, c := range part {
				suit, ok := charToSuit[c]
				if !ok && c == 'z' {
					// honor tile
					for _, v := range values {
						if v < 1 || v > 7 {
							return nil, &InvalidHonorError{}
						}
						suit = Suit(v + 2)
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
								return nil, &RedDoraError{}
							}
						} else {
							newTile = CreateTile(suit, v, false)
						}
						tiles = append(tiles, *newTile)
					}
					values = make([]int, 0)
				}
			}
		} else {
			// melds (pon, chii). kans can be specified as closed; otherwise, they are open.
			values := make([]int, 0)
			for _, c := range part {
				suit, ok := charToSuit[c]
				if !ok && c == 'z' {
					// honor tile
					meldTiles := make([]Tile, 0)
					for _, v := range values {
						if v < 1 || v > 7 {
							return nil, &InvalidHonorError{}
						}
						suit = Suit(v + 2)
						newTile := CreateTile(suit, 0, false)
						tiles = append(tiles, *newTile)
						meldTiles = append(meldTiles, *newTile)

					}
					values = make([]int, 0)
					newMeld, err := CreateMentsu(meldTiles, true)
					if err != nil {
						return nil, err
					}
					melds = append(melds, *newMeld)
				} else if !ok {
					values = append(values, int(c-'0'))
				} else {
					meldTiles := make([]Tile, 0)
					for _, v := range values {
						var newTile *Tile
						if v == 0 {
							// red five
							newTile = CreateTile(suit, 5, true)
							redCounts[suit]++
							if redCounts[suit] > 1 {
								return nil, &RedDoraError{}
							}
						} else {
							newTile = CreateTile(suit, v, false)
						}
						tiles = append(tiles, *newTile)
						meldTiles = append(meldTiles, *newTile)

					}
					values = make([]int, 0)
					newMeld, err := CreateMentsu(meldTiles, true)
					if err != nil {
						return nil, err
					}
					melds = append(melds, *newMeld)
				}
			}
		}
	}
	effectiveTileCount := len(tiles) - kans
	if effectiveTileCount != 13 && effectiveTileCount != 14 {
		return nil, &InvalidHandError{}
	}
	return CreateHand(tiles, melds), nil
}

func tilesToCounts(tiles []Tile) TileCounts {
	var m TileCounts
	for _, tile := range tiles {
		m[TileToID(&tile)]++
	}
	return m
}

func getTileIndex(tiles []Tile, suit Suit, value int) int {
	var l, h int = 0, len(tiles) - 1
	var m int
	for h >= l {
		m = (h + l) / 2
		if tiles[m].value == value && tiles[m].suit == suit {
			return m
		} else if SuitAndValueToID(suit, value) > TileToID(&tiles[m]) {
			l = m + 1
		} else {
			h = m - 1
		}
	}
	return -1
}

func CalculateAllPartitions(h *Hand) []Partition {
	results := make([]Partition, 0)
	nonMeldTiles := make([]Tile, 0)
	// buggy. needs to exclude it from the meld after it is found once.
	for _, tile := range h.tiles {
		nonMeld := true
		for _, meld := range h.melds {
			for _, meldTile := range meld.tiles {
				if meldTile.equals(&tile) {
					nonMeld = false
				}
			}
		}
		if nonMeld {
			nonMeldTiles = append(nonMeldTiles, tile)
		}
	}
	memo := make(map[string][][]Mentsu)
	for _, partition := range calculatePartitionsFromTiles(nonMeldTiles, memo) {
		newPartition := Partition{append(partition, h.melds...)}
		results = append(results, newPartition)
	}
	sort.Slice(results, func(i, j int) bool {
		return len(results[i].Mentsu) < len(results[j].Mentsu)
	})
	return results
}

func calculatePartitionsFromTiles(rest []Tile, memo map[string][][]Mentsu) [][]Mentsu {
	results := make([][]Mentsu, 0)
	if len(rest) == 0 {
		results = append(results, []Mentsu{})
		return results
	}
	key := TilesToString(rest)
	if memoResult, ok := memo[key]; ok {
		return memoResult
	}
	currentTile := rest[0]
	// use first tile as a single
	nextRest := utils.RemoveIndex(utils.Clone(rest), 0)
	singleMentsu, _ := CreateMentsu([]Tile{currentTile}, false)
	for _, partition := range calculatePartitionsFromTiles(nextRest, memo) {
		results = append(results, append(partition, *singleMentsu))
	}
	// try to create a pair
	if index := getTileIndex(nextRest, currentTile.suit, currentTile.value); index >= 0 {
		var pairTile Tile
		results, nextRest, pairTile = removeAndGetPartitions(results, nextRest, index, []Tile{currentTile}, memo)
		// try to create a triplet (no need for quads, as those are melds)
		if index = getTileIndex(nextRest, currentTile.suit, currentTile.value); index >= 0 {
			results, _, _ = removeAndGetPartitions(results, nextRest, index, []Tile{currentTile, pairTile}, memo)
		}
	}
	// try to create a side-wait/sequence
	nextRest = utils.RemoveIndex(utils.Clone(rest), 0)
	if currentTile.value > 0 && currentTile.value < 9 && !currentTile.isHonor() {
		if index := getTileIndex(nextRest, currentTile.suit, currentTile.value+1); index >= 0 {
			var secondTile Tile
			results, nextRest, secondTile = removeAndGetPartitions(results, nextRest, index, []Tile{currentTile}, memo)
			if index := getTileIndex(nextRest, currentTile.suit, currentTile.value+2); currentTile.value < 8 && index >= 0 {
				results, _, _ = removeAndGetPartitions(results, nextRest, index, []Tile{currentTile, secondTile}, memo)
			}
		}
	}
	// try to create a closed-wait
	nextRest = utils.RemoveIndex(utils.Clone(rest), 0)
	if currentTile.value > 0 && currentTile.value < 8 && !currentTile.isHonor() {
		if index := getTileIndex(nextRest, currentTile.suit, currentTile.value+2); index >= 0 {
			results, _, _ = removeAndGetPartitions(results, nextRest, index, []Tile{currentTile}, memo)
		}
	}
	copy(memo[key], results)
	return results
}

func removeAndGetPartitions(results [][]Mentsu, rest []Tile, index int, mentsuTiles []Tile, memo map[string][][]Mentsu) ([][]Mentsu, []Tile, Tile) {
	nextTile := rest[index]
	rest = utils.RemoveIndex(rest, index)
	mentsu, _ := CreateMentsu(append(mentsuTiles, nextTile), false)
	for _, partition := range calculatePartitionsFromTiles(rest, memo) {
		results = append(results, append(partition, *mentsu))
	}
	return results, rest, nextTile
}

func TilesToString(tiles []Tile) string {
	var ret string
	prevSuit, wasHonor := tiles[0].suit, tiles[0].isHonor()
	for _, tile := range tiles {
		if tile.suit != prevSuit && (!tile.isHonor() || !wasHonor) {
			if suit, ok := SuitToString[prevSuit]; ok {
				ret += suit
			} else {
				ret += "z"
			}
			prevSuit = tile.suit
			wasHonor = tile.isHonor()
		}
		if tile.red {
			ret += "5"
		} else if !tile.isHonor() {
			ret += strconv.Itoa(tile.value)
		} else {
			ret += strconv.Itoa(int(tile.suit) - 2)
		}
	}
	if wasHonor {
		ret += "z"
	} else {
		ret += SuitToString[prevSuit]
	}
	return ret
}
