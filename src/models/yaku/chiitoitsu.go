package yaku

import (
	"github.com/mmuiro/riichi-base/src/models"
	"github.com/mmuiro/riichi-base/src/models/constants/languages"
)

type ChiiToitsu struct{}

func (y ChiiToitsu) Match(p *models.Partition, c *Conditions) bool {
	return models.CheckChiiToitsu(p)
}

func (y ChiiToitsu) Han(open bool) int {
	return 2
}

func (y ChiiToitsu) Name(l languages.Language) string {
	if l == languages.EN {
		return "Seven Pairs"
	}
	return "七対子"
}
