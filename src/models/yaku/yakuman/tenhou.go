package yakuman

import (
	"riichi-calculator/src/models"
	"riichi-calculator/src/models/yaku"
)

type Tenhou struct{}

func (y Tenhou) Match(p *models.Partition, c *yaku.Conditions) bool {
	return c.Tenhou
}

func (y Tenhou) Value(open bool) int {
	return 1
}

func (y Tenhou) Description() string {
	return "Complete hand at dealer's first turn."
}

func (y Tenhou) Name() string {
	return "Tenhou"
}
