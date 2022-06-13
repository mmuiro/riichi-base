package yakuman

import (
	"riichi-calculator/src/models"
	"riichi-calculator/src/models/constants/waits"
	"riichi-calculator/src/models/yaku"
)

type KokushiThirteen struct{}

func (y KokushiThirteen) Match(p *models.Partition, c *yaku.Conditions) bool {
	return models.CheckKokushi(p) && p.Wait == waits.KokushiThirteen
}

func (y KokushiThirteen) Value() int {
	return 2
}

func (y KokushiThirteen) Description() string {
	return "Thirteen orphans, thirteen-way wait."
}

func (y KokushiThirteen) Name() string {
	return "Kokushi Musou Juusan Men Machi"
}
