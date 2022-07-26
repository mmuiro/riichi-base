package yaku

import (
	"github.com/mmuiro/riichi-base/src/models"
	"github.com/mmuiro/riichi-base/src/models/constants/groups"
	"github.com/mmuiro/riichi-base/src/models/constants/languages"
)

type YakuhaiBakaze struct{}

func (y YakuhaiBakaze) Match(p *models.Partition, c *Conditions) bool {
	for _, mentsu := range p.Mentsu {
		if mentsu.Kind == groups.Kanchan || mentsu.Kind == groups.Koutsu && mentsu.Suit == c.Bakaze {
			return true
		}
	}
	return false
}

func (y YakuhaiBakaze) Han(open bool) int { return 1 }

func (y YakuhaiBakaze) Name(l languages.Language) string {
	if l == languages.EN {
		return "Round Wind"
	}
	return "役牌：場風牌"
}
