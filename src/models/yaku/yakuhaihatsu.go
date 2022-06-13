package yaku

import (
	"riichi-calculator/src/models"
	"riichi-calculator/src/models/constants/groups"
	"riichi-calculator/src/models/constants/suits"
)

type YakuhaiHatsu struct{}

func (y YakuhaiHatsu) Match(p *models.Partition, c *Conditions) bool {
	for _, mentsu := range p.Mentsu {
		if mentsu.Kind == groups.Kanchan || mentsu.Kind == groups.Koutsu && mentsu.Suit == suits.Hatsu {
			return true
		}
	}
	return false
}

func (y YakuhaiHatsu) Han(open bool) int { return 1 }

func (y YakuhaiHatsu) Description() string {
	return "Set of Hatsu."
}

func (y YakuhaiHatsu) Name() string {
	return "Yakuhai: Hatsu"
}
