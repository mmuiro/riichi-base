package yaku

import (
	"github.com/mmuiro/riichi-base/src/models"
	"github.com/mmuiro/riichi-base/src/models/constants/groups"
	"github.com/mmuiro/riichi-base/src/models/constants/languages"
	"github.com/mmuiro/riichi-base/src/models/constants/suits"
)

type YakuhaiHaku struct{}

func (y YakuhaiHaku) Match(p *models.Partition, c *Conditions) bool {
	for _, mentsu := range p.Mentsu {
		if mentsu.Kind == groups.Kanchan || mentsu.Kind == groups.Koutsu && mentsu.Suit == suits.Haku {
			return true
		}
	}
	return false
}

func (y YakuhaiHaku) Han(open bool) int { return 1 }

func (y YakuhaiHaku) Name(l languages.Language) string {
	if l == languages.EN {
		return "White Dragon"
	}
	return "役牌白"
}
