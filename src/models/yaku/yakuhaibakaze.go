package yaku

import (
	"riichi-calculator/src/models"
	"riichi-calculator/src/models/constants/groups"
)

type YakuhaiBakaze struct{}

func (y *YakuhaiBakaze) Match(p *models.Partition, c *Conditions) bool {
	for _, mentsu := range p.Mentsu {
		if mentsu.Kind == groups.Kanchan || mentsu.Kind == groups.Koutsu && mentsu.Suit == c.Bakaze {
			return true
		}
	}
	return false
}

func (y *YakuhaiBakaze) Han(open bool) int { return 1 }

func (y *YakuhaiBakaze) Description() string {
	return "Set of Round Wind."
}
