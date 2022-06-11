package models

import (
	"riichi-calculator/src/models/constants/groups"
	"riichi-calculator/src/models/constants/suits"
	"sort"
	"strings"
)

type Mentsu struct {
	tiles    []Tile
	open     bool
	complete bool
	suit     suits.Suit
	kind     groups.MentsuType
}

func (m Mentsu) String() string {
	tileStrings := make([]string, len(m.tiles))
	for i, tile := range m.tiles {
		tileStrings[i] = tile.String()
	}
	ret := strings.Join(tileStrings, " ")
	return ret
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
	mentsu := &Mentsu{tiles: tiles, open: open}
	if check, err := checkAndAssignKind(mentsu); !check {
		return nil, err
	}
	mentsu.suit = tiles[0].Suit
	return mentsu, nil
}

func (m *Mentsu) addTile(t *Tile) error {
	m.tiles = append(m.tiles, *t)
	sort.Slice(m.tiles, func(i, j int) bool {
		return TileToID(&m.tiles[i]) < TileToID(&m.tiles[j])
	})
	if check, err := checkAndAssignKind(m); !check {
		return err
	}
	return nil
}

// Returns if the given group is a triplet.
func checkKoutsu(m *Mentsu) bool {
	return len(m.tiles) == 3 && m.tiles[0].Equals(&m.tiles[1]) && m.tiles[0].Equals(&m.tiles[2])
}

// Returns if the given group is a quad.
func checkKantsu(m *Mentsu) bool {
	return len(m.tiles) == 4 && m.tiles[0].Equals(&m.tiles[1]) && m.tiles[0].Equals(&m.tiles[2]) && m.tiles[0].Equals(&m.tiles[3])
}

// Returns if the given group is a pair.
func checkToitsu(m *Mentsu) bool {
	return len(m.tiles) == 2 && m.tiles[0].Equals(&m.tiles[1])
}

// Returns if the given group is a sequence. Assumes tiles are of the same suit.
func checkShuntsu(m *Mentsu) bool {
	return len(m.tiles) == 3 && m.tiles[1].Value == m.tiles[0].Value+1 && m.tiles[2].Value == m.tiles[1].Value+1
}

// Returns if the given group is a ryanmen(side) wait proto-sequence.
func checkRyanmen(m *Mentsu) bool {
	return len(m.tiles) == 2 && m.tiles[1].Value == m.tiles[0].Value+1 && 1 < m.tiles[0].Value && m.tiles[0].Value < 8
}

// Returns if the given group is a kanchan(closed) wait proto-sequence.
func checkKanchan(m *Mentsu) bool {
	return len(m.tiles) == 2 && m.tiles[1].Value == m.tiles[0].Value+2
}

// Returns if the given group is a penchan(edge) wait proto-sequence.
func checkPenchan(m *Mentsu) bool {
	return len(m.tiles) == 2 && m.tiles[1].Value == m.tiles[0].Value+1 && (m.tiles[0].Value == 1 || m.tiles[1].Value == 9)
}

// Returns if the given group is a single tile.
func checkSingle(m *Mentsu) bool {
	return len(m.tiles) == 1
}

// Returns if the given group is valid. If so, also assigns the group kind to the group.
func checkAndAssignKind(m *Mentsu) (bool, error) {
	if len(m.tiles) > 4 {
		return false, &InvalidGroupError{}
	}
	unique_suits := make(map[suits.Suit]bool)
	for _, tile := range m.tiles {
		unique_suits[tile.Suit] = true
	}
	if len(unique_suits) != 1 {
		return false, &MultipleSuitError{}
	}
	checks := [](func(m *Mentsu) bool){checkToitsu, checkShuntsu, checkKoutsu, checkKantsu, checkSingle, checkRyanmen, checkKanchan, checkPenchan}
	for i, f := range checks {
		if f(m) {
			m.kind = groups.MentsuType(i)
			return true, nil
		}
	}
	return false, &InvalidGroupError{}
}
