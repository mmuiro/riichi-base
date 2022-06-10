package models

import (
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type TileCounts [34]int

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

var charToSuit = map[rune]SuitType{
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
	kans, redCounts := 0, make(map[SuitType]int)
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

func parsePartToTiles(part string, redCounts map[SuitType]int, defaultOpen bool) ([]Tile, bool, error) {
	tiles := make([]Tile, 0)
	values := make([]int, 0)
	open := defaultOpen
	for _, c := range part {
		suit, ok := charToSuit[c]
		if !ok && c == 'z' {
			// honor tile
			for _, v := range values {
				if v < 1 || v > 7 {
					return nil, open, &InvalidHonorError{}
				}
				suit = SuitType(v + 2)
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

func getTileIndex(tiles []Tile, suit SuitType, value int) int {
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
			if suit, ok := SuitToString[prevSuit]; ok {
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
		ret += SuitToString[prevSuit]
	}
	return ret
}
