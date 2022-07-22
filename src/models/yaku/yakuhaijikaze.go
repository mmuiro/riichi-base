package yaku

import (
	"github.com/mmuiro/riichi-base/src/models"
	"github.com/mmuiro/riichi-base/src/models/constants/groups"
)

type YakuhaiJikaze struct{}

func (y YakuhaiJikaze) Match(p *models.Partition, c *Conditions) bool {
	for _, mentsu := range p.Mentsu {
		if mentsu.Kind == groups.Kanchan || mentsu.Kind == groups.Koutsu && mentsu.Suit == c.Jikaze {
			return true
		}
	}
	return false
}

func (y YakuhaiJikaze) Han(open bool) int { return 1 }

func (y YakuhaiJikaze) Description() string {
	return "Set of Round Wind."
}

func (y YakuhaiJikaze) Name() string {
	return "Yakuhai: Seat Wind"
}
