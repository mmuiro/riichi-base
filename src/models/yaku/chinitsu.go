package yaku

import (
	"riichi-calculator/src/models"
	"riichi-calculator/src/utils"
)

type Chinitsu struct{}

func (y Chinitsu) Match(p *models.Partition, c *Conditions) bool {
	return utils.All(utils.FuncMap(func(m models.Mentsu) bool {
		return !m.Tiles[0].IsHonor() && m.Suit == p.Mentsu[0].Suit
	}, p.Mentsu))
}

func (y Chinitsu) Han(open bool) int {
	if open {
		return 5
	}
	return 6
}

func (y Chinitsu) Description() string {
	return "Full flush of a suit."
}

func (y Chinitsu) Name() string {
	return "Chinitsu"
}
