package models

import (
	"riichi-calculator/src/models/constants/groups"
	"riichi-calculator/src/models/constants/suits"
	"sort"
	"strings"
)

type Mentsu struct {
	Tiles []Tile
	Open  bool
	Suit  suits.Suit
	Kind  groups.MentsuType
}

func (m Mentsu) String() string {
	tileStrings := make([]string, len(m.Tiles))
	for i, tile := range m.Tiles {
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
	mentsu := &Mentsu{Tiles: tiles, Open: open}
	if check, err := checkAndAssignKind(mentsu); !check {
		return nil, err
	}
	mentsu.Suit = tiles[0].Suit
	return mentsu, nil
}

func (m *Mentsu) addTile(t *Tile) error {
	m.Tiles = append(m.Tiles, *t)
	sort.Slice(m.Tiles, func(i, j int) bool {
		return TileToID(&m.Tiles[i]) < TileToID(&m.Tiles[j])
	})
	if check, err := checkAndAssignKind(m); !check {
		return err
	}
	return nil
}

// Returns if the given group is a triplet.
func checkKoutsu(m *Mentsu) bool {
	return len(m.Tiles) == 3 && m.Tiles[0].Equals(&m.Tiles[1]) && m.Tiles[0].Equals(&m.Tiles[2])
}

// Returns if the given group is a quad.
func checkKantsu(m *Mentsu) bool {
	return len(m.Tiles) == 4 && m.Tiles[0].Equals(&m.Tiles[1]) && m.Tiles[0].Equals(&m.Tiles[2]) && m.Tiles[0].Equals(&m.Tiles[3])
}

// Returns if the given group is a pair.
func checkToitsu(m *Mentsu) bool {
	return len(m.Tiles) == 2 && m.Tiles[0].Equals(&m.Tiles[1])
}

// Returns if the given group is a sequence. Assumes tiles are of the same suit.
func checkShuntsu(m *Mentsu) bool {
	return len(m.Tiles) == 3 && m.Tiles[1].Value == m.Tiles[0].Value+1 && m.Tiles[2].Value == m.Tiles[1].Value+1
}

// Returns if the given group is a ryanmen(side) wait proto-sequence.
func checkRyanmen(m *Mentsu) bool {
	return len(m.Tiles) == 2 && m.Tiles[1].Value == m.Tiles[0].Value+1 && 1 < m.Tiles[0].Value && m.Tiles[0].Value < 8
}

// Returns if the given group is a kanchan(closed) wait proto-sequence.
func checkKanchan(m *Mentsu) bool {
	return len(m.Tiles) == 2 && m.Tiles[1].Value == m.Tiles[0].Value+2
}

// Returns if the given group is a penchan(edge) wait proto-sequence.
func checkPenchan(m *Mentsu) bool {
	return len(m.Tiles) == 2 && m.Tiles[1].Value == m.Tiles[0].Value+1 && (m.Tiles[0].Value == 1 || m.Tiles[1].Value == 9)
}

// Returns if the given group is a single tile.
func checkSingle(m *Mentsu) bool {
	return len(m.Tiles) == 1
}

// Returns if the given group is valid. If so, also assigns the group kind to the group.
func checkAndAssignKind(m *Mentsu) (bool, error) {
	if len(m.Tiles) > 4 {
		return false, &InvalidGroupError{}
	}
	unique_suits := make(map[suits.Suit]bool)
	for _, tile := range m.Tiles {
		unique_suits[tile.Suit] = true
	}
	if len(unique_suits) != 1 {
		return false, &MultipleSuitError{}
	}
	checks := [](func(m *Mentsu) bool){checkToitsu, checkShuntsu, checkKoutsu, checkKantsu, checkSingle, checkRyanmen, checkKanchan, checkPenchan}
	for i, f := range checks {
		if f(m) {
			m.Kind = groups.MentsuType(i)
			return true, nil
		}
	}
	return false, &InvalidGroupError{}
}
