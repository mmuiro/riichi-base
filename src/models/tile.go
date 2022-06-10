package models

type SuitType int

const (
	Man SuitType = iota
	Sou
	Pin
	Ton
	Nan
	Xia
	Pei
	Haku
	Hatsu
	Chun
)

var TileToString = [34]string{
	"🀇", "🀈", "🀉", "🀊", "🀋", "🀌", "🀍", "🀎", "🀏",
	"🀐", "🀑", "🀒", "🀓", "🀔", "🀕", "🀖", "🀗", "🀘",
	"🀙", "🀚", "🀛", "🀜", "🀝", "🀞", "🀟", "🀠", "🀡",
	"🀀", "🀁", "🀂", "🀃", "🀆", "🀅", "🀄",
}

var SuitToString = map[SuitType]string{
	Man: "m",
	Sou: "s",
	Pin: "p",
}

type Tile struct {
	Suit  SuitType
	Value int
	Red   bool
}

func TileToID(t *Tile) int {
	return SuitAndValueToID(t.Suit, t.Value)
}

func SuitAndValueToID(suit SuitType, value int) int {
	if suit < Ton {
		return (int(suit))*9 + value - 1
	} else {
		return int(suit) + 24
	}
}

func CreateTile(suit SuitType, value int, red bool) *Tile {
	t := &Tile{Suit: suit, Value: value, Red: red}
	return t
}

func (t Tile) IsHonor() bool {
	return t.Suit >= Ton && t.Suit <= Chun
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
