package yaku

import "github.com/mmuiro/riichi-base/src/models/constants/suits"

type Conditions struct {
	Tsumo        bool
	Riichi       bool
	DoubleRiichi bool
	Ippatsu      bool
	Menzenchin   bool
	Houtei       bool
	Haitei       bool
	Rinshan      bool
	Chankan      bool
	Tenhou       bool
	Chiihou      bool
	Bakaze       suits.Suit
	Jikaze       suits.Suit
	Dora         []int
	UraDora      []int
}

// https://riichi.wiki/List_of_terminology_by_alphabetical_order
