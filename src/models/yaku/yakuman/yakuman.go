package yakuman

import (
	"github.com/mmuiro/riichi-base/src/models"
	"github.com/mmuiro/riichi-base/src/models/constants/languages"
	"github.com/mmuiro/riichi-base/src/models/yaku"
)

type Yakuman interface {
	Match(p *models.Partition, c *yaku.Conditions) bool
	Value() int
	Name(l languages.Language) string
}

var AllYakuman = []Yakuman{
	Kokushi{},
	KokushiThirteen{},
	SuuAnkou{},
	SuuAnkouTanki{},
	DaiSangen{},
	ShouSuushii{},
	DaiSuushii{},
	TsuuIisou{},
	Chinroutou{},
	RyuuIisou{},
	Chuuren{},
	JunseiChuuren{},
	SuuKantsu{},
	Tenhou{},
	Chiihou{},
}
