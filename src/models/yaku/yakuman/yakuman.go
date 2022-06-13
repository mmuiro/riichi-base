package yakuman

import (
	"riichi-calculator/src/models"
	"riichi-calculator/src/models/yaku"
)

type Yakuman interface {
	Match(p *models.Partition, c *yaku.Conditions) bool
	Value() int
	Name() string
	Description() string
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
