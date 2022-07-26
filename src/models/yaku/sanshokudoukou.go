package yaku

import (
	"github.com/mmuiro/riichi-base/src/models"
	"github.com/mmuiro/riichi-base/src/models/constants/groups"
	"github.com/mmuiro/riichi-base/src/models/constants/languages"
)

type SanshokuDoukou struct{}

func (y SanshokuDoukou) Match(p *models.Partition, c *Conditions) bool {
	uniqueKoutsu := make(map[int]bool)
	for _, mentsu := range p.Mentsu {
		if mentsu.Kind == groups.Koutsu && !mentsu.Tiles[0].IsHonor() {
			id := models.TileToID(&mentsu.Tiles[0])
			uniqueKoutsu[id] = true
			if uniqueKoutsu[id%9] && uniqueKoutsu[(id%9)+9] && uniqueKoutsu[(id%9)+18] {
				return true
			}
		}
	}
	return false
}

func (y SanshokuDoukou) Han(open bool) int {
	return 2
}

func (y SanshokuDoukou) Name(l languages.Language) string {
	if l == languages.EN {
		return "Triple Triplets"
	}
	return "三色同刻"
}
