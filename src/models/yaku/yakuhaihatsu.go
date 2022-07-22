package yaku

import (
	"github.com/mmuiro/riichi-base/src/models"
	"github.com/mmuiro/riichi-base/src/models/constants/groups"
	"github.com/mmuiro/riichi-base/src/models/constants/suits"
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
