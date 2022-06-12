package yaku

import (
	"riichi-calculator/src/models"
	"riichi-calculator/src/models/constants/groups"
	"riichi-calculator/src/models/constants/suits"
)

type YakuhaiHaku struct{}

func (y *YakuhaiHaku) Match(p *models.Partition, c *Conditions) bool {
	for _, mentsu := range p.Mentsu {
		if mentsu.Kind == groups.Kanchan || mentsu.Kind == groups.Koutsu && mentsu.Suit == suits.Haku {
			return true
		}
	}
	return false
}

func (y *YakuhaiHaku) Han(open bool) int { return 1 }

func (y *YakuhaiHaku) Description() string {
	return "Set of Haku."
}
