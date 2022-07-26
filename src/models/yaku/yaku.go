package yaku

import (
	"github.com/mmuiro/riichi-base/src/models"
	"github.com/mmuiro/riichi-base/src/models/constants/languages"
)

type Yaku interface {
	Match(p *models.Partition, c *Conditions) bool
	Han(open bool) int
	Name(l languages.Language) string
}

var AllYaku = []Yaku{
	MenzenTsumo{},
	Riichi{},
	Ippatsu{},
	Pinfu{},
	Iipeikou{},
	Haitei{},
	Houtei{},
	RinshanKaihou{},
	Chankan{},
	Tanyao{},
	YakuhaiBakaze{},
	YakuhaiJikaze{},
	YakuhaiChun{},
	YakuhaiHaku{},
	YakuhaiHatsu{},
	DoubleRiichi{},
	Chanta{},
	SanshokuDoujun{},
	Ittsu{},
	Toitoi{},
	SanAnkou{},
	SanshokuDoukou{},
	SanKantsu{},
	ChiiToitsu{},
	HonRoutou{},
	ShouSangen{},
	Honitsu{},
	JunchanTaiyao{},
	Ryanpeikou{},
	Chinitsu{},
}
