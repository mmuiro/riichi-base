package yaku

import (
	"riichi-calculator/src/models"
	"riichi-calculator/src/models/constants/groups"
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

func (y SanshokuDoujun) Description() string {
	return "Two identical sequences on a closed hand."
}

func (y SanshokuDoujun) Name() string {
	return "Sanshoku Doujun"
}
