package yaku

import (
	"github.com/mmuiro/riichi-base/src/models"
	"github.com/mmuiro/riichi-base/src/models/constants/groups"
)

type Iipeikou struct{}

func (y Iipeikou) Match(p *models.Partition, c *Conditions) bool {
	if !c.Menzenchin {
		return false
	}
	uniqueShuntsu := make(map[int]int)
	for _, mentsu := range p.Mentsu {
		if mentsu.Kind == groups.Shuntsu {
			uniqueShuntsu[models.TileToID(&mentsu.Tiles[0])]++
		}
	}
	multi := 0
	for _, count := range uniqueShuntsu {
		if count >= 2 { // we don't have to worry about the four case - it's suu ankou
			multi++
		}
	}
	return multi == 1
}

func (y Iipeikou) Han(open bool) int { return 1 }

func (y Iipeikou) Description() string {
	return "Two identical sequences on a closed hand."
}

func (y Iipeikou) Name() string {
	return "Iipeikou"
}
