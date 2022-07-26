package yaku

import (
	"github.com/mmuiro/riichi-base/src/models"
	"github.com/mmuiro/riichi-base/src/models/constants/groups"
	"github.com/mmuiro/riichi-base/src/models/constants/languages"
	"github.com/mmuiro/riichi-base/src/models/constants/suits"
)

type Ittsu struct{}

func (y Ittsu) Match(p *models.Partition, c *Conditions) bool {
	uniqueShuntsu := make(map[int]bool)
	for _, mentsu := range p.Mentsu {
		if mentsu.Kind == groups.Shuntsu {
			uniqueShuntsu[models.TileToID(&mentsu.Tiles[0])] = true
		}
	}
	for suit := suits.Man; suit <= suits.Pin; suit++ {
		if uniqueShuntsu[int(suit)*9] && uniqueShuntsu[int(suit)*9+3] && uniqueShuntsu[int(suit)*9+6] {
			return true
		}
	}
	return false
}

func (y Ittsu) Han(open bool) int {
	if open {
		return 1
	}
	return 2
}

func (y Ittsu) Name(l languages.Language) string {
	if l == languages.EN {
		return "Pure Straight"
	}
	return "一気通貫"
}
