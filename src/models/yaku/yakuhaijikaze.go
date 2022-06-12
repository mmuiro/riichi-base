package yaku

import (
	"riichi-calculator/src/models"
	"riichi-calculator/src/models/constants/groups"
)

type YakuhaiJikaze struct{}

func (y *YakuhaiJikaze) Match(p *models.Partition, c *Conditions) bool {
	for _, mentsu := range p.Mentsu {
		if mentsu.Kind == groups.Kanchan || mentsu.Kind == groups.Koutsu && mentsu.Suit == c.Jikaze {
			return true
		}
	}
	return false
}

func (y *YakuhaiJikaze) Han() int { return 1 }

func (y *YakuhaiJikaze) Description() string {
	return "Set of Round Wind."
}
