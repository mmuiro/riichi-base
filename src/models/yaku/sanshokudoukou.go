package yaku

import (
	"riichi-calculator/src/models"
	"riichi-calculator/src/models/constants/groups"
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

func (y SanshokuDoukou) Description() string {
	return "Three mixed-suits triplets of the same number."
}

func (y SanshokuDoukou) Name() string {
	return "Sanshoku Doukou"
}
