package yaku

import (
	"riichi-calculator/src/models"
	"riichi-calculator/src/models/constants/groups"
	"riichi-calculator/src/models/constants/suits"
)

type YakuhaiChun struct{}

func (y YakuhaiChun) Match(p *models.Partition, c *Conditions) bool {
	for _, mentsu := range p.Mentsu {
		if mentsu.Kind == groups.Kanchan || mentsu.Kind == groups.Koutsu && mentsu.Suit == suits.Chun {
			return true
		}
	}
	return false
}

func (y YakuhaiChun) Han(open bool) int { return 1 }

func (y YakuhaiChun) Description() string {
	return "Set of Chun."
}

func (y YakuhaiChun) Name() string {
	return "Yakuhai: Chun"
}
