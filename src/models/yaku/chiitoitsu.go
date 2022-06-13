package yaku

import (
	"riichi-calculator/src/models"
)

type ChiiToitsu struct{}

func (y ChiiToitsu) Match(p *models.Partition, c *Conditions) bool {
	return models.CheckChiiToitsu(p)
}

func (y ChiiToitsu) Han(open bool) int {
	return 2
}

func (y ChiiToitsu) Description() string {
	return "All pairs."
}

func (y ChiiToitsu) Name() string {
	return "Chii Toitsu"
}
