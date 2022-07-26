package yaku

import (
	"github.com/mmuiro/riichi-base/src/models"
	"github.com/mmuiro/riichi-base/src/models/constants/groups"
	"github.com/mmuiro/riichi-base/src/models/constants/languages"
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

func (y YakuhaiJikaze) Name(l languages.Language) string {
	if l == languages.EN {
		return "Seat Wind"
	}
	return "役牌：自風牌"
}
