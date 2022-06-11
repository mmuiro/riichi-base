package suits

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

var SuitToString = map[Suit]string{
	Man: "m",
	Sou: "s",
	Pin: "p",
}

var CharToSuit = map[rune]Suit{
	'm': Man,
	's': Sou,
	'p': Pin,
}
