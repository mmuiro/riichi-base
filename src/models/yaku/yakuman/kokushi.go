package yakuman

import (
	"riichi-calculator/src/models"
	"riichi-calculator/src/models/constants/waits"
	"riichi-calculator/src/models/yaku"
)

type Kokushi struct{}

func (y Kokushi) Match(p *models.Partition, c *yaku.Conditions) bool {
	return models.CheckKokushi(p) && p.Wait == waits.KokushiSingle
}

func (y Kokushi) Value(open bool) int {
	return 1
}

func (y Kokushi) Description() string {
	return "Thirteen orphans."
}

func (y Kokushi) Name() string {
	return "Kokushi Musou"
}
