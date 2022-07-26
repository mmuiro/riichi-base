package yaku

import (
	"github.com/mmuiro/riichi-base/src/models"
	"github.com/mmuiro/riichi-base/src/models/constants/groups"
	"github.com/mmuiro/riichi-base/src/models/constants/languages"
)

type SanshokuDoujun struct{}

func (y SanshokuDoujun) Match(p *models.Partition, c *Conditions) bool {
	uniqueShuntsu := make(map[int]bool)
	for _, mentsu := range p.Mentsu {
		if mentsu.Kind == groups.Shuntsu {
			id := models.TileToID(&mentsu.Tiles[0])
			uniqueShuntsu[id] = true
			if uniqueShuntsu[id%9] && uniqueShuntsu[(id%9)+9] && uniqueShuntsu[(id%9)+18] {
				return true
			}
		}
	}
	return false
}

func (y SanshokuDoujun) Han(open bool) int {
	if open {
		return 1
	}
	return 2
}

func (y SanshokuDoujun) Name(l languages.Language) string {
	if l == languages.EN {
		return "Triple Mixed Sequence"
	}
	return "三色同順"
}
