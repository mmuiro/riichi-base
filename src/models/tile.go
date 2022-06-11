package models

import "riichi-calculator/src/models/constants/suits"

var TileToString = [34]string{
	"🀇", "🀈", "🀉", "🀊", "🀋", "🀌", "🀍", "🀎", "🀏",
	"🀐", "🀑", "🀒", "🀓", "🀔", "🀕", "🀖", "🀗", "🀘",
	"🀙", "🀚", "🀛", "🀜", "🀝", "🀞", "🀟", "🀠", "🀡",
	"🀀", "🀁", "🀂", "🀃", "🀆", "🀅", "🀄",
}

type Tile struct {
	Suit  suits.Suit
	Value int
	Red   bool
}

func TileToID(t *Tile) int {
	return SuitAndValueToID(t.Suit, t.Value)
}

func SuitAndValueToID(suit suits.Suit, value int) int {
	if suit < suits.Ton {
		return (int(suit))*9 + value - 1
	} else {
		return int(suit) + 24
	}
}

func CreateTile(suit suits.Suit, value int, red bool) *Tile {
	t := &Tile{Suit: suit, Value: value, Red: red}
	return t
}

func (t Tile) IsHonor() bool {
	return t.Suit >= suits.Ton && t.Suit <= suits.Chun
}

func (t Tile) Equals(other *Tile) bool {
	suit_check := t.Suit == other.Suit
	if t.IsHonor() {
		return suit_check
	}
	return suit_check && t.Value == other.Value
}

func (t Tile) String() string {
	return TileToString[TileToID(&t)]
}
