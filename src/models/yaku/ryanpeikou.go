package yaku

import (
	"github.com/mmuiro/riichi-base/src/models"
	"github.com/mmuiro/riichi-base/src/models/constants/groups"
	"github.com/mmuiro/riichi-base/src/models/constants/languages"
)

type Ryanpeikou struct{}

func (y Ryanpeikou) Match(p *models.Partition, c *Conditions) bool {
	if !c.Menzenchin {
		return false
	}
	uniqueShuntsu := make(map[int]int)
	for _, mentsu := range p.Mentsu {
		if mentsu.Kind == groups.Shuntsu {
			uniqueShuntsu[models.TileToID(&mentsu.Tiles[0])]++
		}
	}
	for _, count := range uniqueShuntsu {
		if count != 2 {
			return false
		}
	}
	return len(uniqueShuntsu) == 2
}

func (y Ryanpeikou) Han(open bool) int { return 3 }

func (y Ryanpeikou) Name(l languages.Language) string {
	if l == languages.EN {
		return "Twice Pure Double Sequence"
	}
	return "二盃口"
}
