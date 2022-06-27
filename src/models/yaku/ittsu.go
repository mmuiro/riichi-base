package yaku

import (
	"riichi-calculator/src/models"
	"riichi-calculator/src/models/constants/groups"
	"riichi-calculator/src/models/constants/suits"
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

func (y Ittsu) Description() string {
	return "A full set of 3 sequences from 1 to 9 of a single suit."
}

func (y Ittsu) Name() string {
	return "Ittsu"
}
