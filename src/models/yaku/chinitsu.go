package yaku

import (
	"github.com/mmuiro/riichi-base/src/models"
	"github.com/mmuiro/riichi-base/src/models/constants/languages"
	"github.com/mmuiro/riichi-base/src/utils"
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

func (y Chinitsu) Name(l languages.Language) string {
	if l == languages.EN {
		return "Full Flush"
	}
	return "清一色"
}
