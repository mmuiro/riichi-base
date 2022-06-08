package models

import (
	"sort"
	"strings"
)

type Mentsu struct {
	tiles    []Tile
	open     bool
	complete bool
	suit     Suit
}

func (m Mentsu) String() string {
	tileStrings := make([]string, len(m.tiles))
	for i, tile := range m.tiles {
		tileStrings[i] = tile.String()
	}
	return strings.Join(tileStrings, " ")
}

type MultipleSuitError struct{}

func (m *MultipleSuitError) Error() string {
	return "A mentsu cannot have multiple suits."
}

type InvalidGroupError struct{}

func (m *InvalidGroupError) Error() string {
	return "A mentsu must be a proto-sequence, single, pair, or set."
}

// Creates a group with the given tiles, if they form a valid one.
func CreateMentsu(tiles []Tile, open bool) (*Mentsu, error) {
	sort.Slice(tiles, func(i, j int) bool {
		return tiles[i].suit < tiles[j].suit || tiles[i].value < tiles[j].value
	}) // leave the sorting to the thing that provides tiles?
	mentsu := &Mentsu{tiles: tiles, open: open}
	if check, err := checkValid(mentsu); !check {
		return nil, err
	}
	mentsu.suit = tiles[0].suit
	return mentsu, nil
}

// Returns if the given group is a set (triplet or quad).
func checkSet(m *Mentsu) bool {
	if len(m.tiles) == 3 {
		return m.tiles[0].equals(&m.tiles[1]) && m.tiles[0].equals(&m.tiles[2])
	} else if len(m.tiles) == 4 {
		return m.tiles[0].equals(&m.tiles[1]) && m.tiles[0].equals(&m.tiles[2]) && m.tiles[0].equals(&m.tiles[3])
	}
	return false
}

// Returns if the given group is a pair.
func checkPair(m *Mentsu) bool {
	return len(m.tiles) == 2 && m.tiles[0].equals(&m.tiles[1])
}

// Returns if the given group is a sequence. Assumes tiles are of the same suit.
func checkSequence(m *Mentsu) bool {
	return len(m.tiles) == 3 && m.tiles[1].value == m.tiles[0].value+1 && m.tiles[2].value == m.tiles[1].value+1
}

func checkProtoSequence(m *Mentsu) bool {
	return len(m.tiles) == 2 && m.tiles[1].value == m.tiles[0].value+2 || m.tiles[1].value == m.tiles[0].value+1
}

func checkSingle(m *Mentsu) bool {
	return len(m.tiles) == 1
}

// Returns if the given group is valid.
func checkValid(m *Mentsu) (bool, error) {
	if len(m.tiles) > 4 {
		return false, &InvalidGroupError{}
	}
	unique_suits := make(map[Suit]bool)
	for _, tile := range m.tiles {
		unique_suits[tile.suit] = true
	}
	if len(unique_suits) != 1 {
		return false, &MultipleSuitError{}
	}
	if checkPair(m) || checkSequence(m) || checkSet(m) {
		return true, nil
	} else if checkSingle(m) || checkProtoSequence(m) {
		return true, nil
	}
	return false, &InvalidGroupError{}
}
