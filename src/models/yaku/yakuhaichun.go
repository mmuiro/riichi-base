package yaku

import (
	"github.com/mmuiro/riichi-base/src/models"
	"github.com/mmuiro/riichi-base/src/models/constants/groups"
	"github.com/mmuiro/riichi-base/src/models/constants/languages"
	"github.com/mmuiro/riichi-base/src/models/constants/suits"
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

func (y YakuhaiChun) Name(l languages.Language) string {
	if l == languages.EN {
		return "Red Dragon"
	}
	return "役牌中"
}
