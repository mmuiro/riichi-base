package yaku

import (
	"riichi-calculator/src/models"
	"riichi-calculator/src/models/constants/groups"
)

type Iipeikou struct{}

func (y *Iipeikou) Match(p *models.Partition, c *Conditions) bool {
	if !c.Menzenchin {
		return false
	}
	uniqueShuntsu := make(map[int]bool)
	for _, mentsu := range p.Mentsu {
		if mentsu.Kind == groups.Shuntsu {
			if uniqueShuntsu[models.TileToID(&mentsu.Tiles[0])] {
				return true
			}
			uniqueShuntsu[models.TileToID(&mentsu.Tiles[0])] = true
		}
	}
	return false
}

func (y *Iipeikou) Han(open bool) int { return 1 }

func (y *Iipeikou) Description() string {
	return "Two identical sequences on a closed hand."
}
