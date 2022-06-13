package yaku

import (
	"riichi-calculator/src/models"
	"riichi-calculator/src/models/constants/suits"
)

type Honitsu struct{}

func (y Honitsu) Match(p *models.Partition, c *Conditions) bool {
	flushSuit := suits.Suit(-1)
	hasHonor := false
	for _, mentsu := range p.Mentsu {
		if !mentsu.Tiles[0].IsHonor() {
			if flushSuit == -1 {
				flushSuit = mentsu.Suit
			} else if mentsu.Suit != flushSuit {
				return false
			}
		} else {
			hasHonor = true
		}
	}
	return hasHonor
}

func (y Honitsu) Han(open bool) int {
	if open {
		return 2
	}
	return 3
}

func (y Honitsu) Description() string {
	return "Half flush of a suit."
}

func (y Honitsu) Name() string {
	return "Honitsu"
}
