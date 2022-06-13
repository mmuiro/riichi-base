package yakuman

import (
	"riichi-calculator/src/models"
	"riichi-calculator/src/models/yaku"
)

type Chiihou struct{}

func (y Chiihou) Match(p *models.Partition, c *yaku.Conditions) bool {
	return c.Chiihou
}

func (y Chiihou) Value() int {
	return 1
}

func (y Chiihou) Description() string {
	return "Complete hand at non-dealer's first draw, before any calls."
}

func (y Chiihou) Name() string {
	return "Chiihou"
}
