package models

type Suit int

const (
	Man Suit = iota
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

var SuitToString = map[Suit]string{
	Man: "m",
	Sou: "s",
	Pin: "p",
}

type Tile struct {
	suit  Suit
	value int
	red   bool
}

func TileToID(t *Tile) int {
	return SuitAndValueToID(t.suit, t.value)
}

func SuitAndValueToID(suit Suit, value int) int {
	if suit < Ton {
		return (int(suit))*9 + value - 1
	} else {
		return int(suit) + 24
	}
}

func CreateTile(suit Suit, value int, red bool) *Tile {
	t := &Tile{suit: suit, value: value, red: red}
	return t
}

func (t Tile) isHonor() bool {
	return t.suit >= Ton && t.suit <= Chun
}

func (t Tile) equals(other *Tile) bool {
	suit_check := t.suit == other.suit
	if t.isHonor() {
		return suit_check
	}
	return suit_check && t.value == other.value
}

func (t Tile) String() string {
	return TileToString[TileToID(&t)]
}
