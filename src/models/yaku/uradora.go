package yaku

import (
	"github.com/mmuiro/riichi-base/src/models"
	"github.com/mmuiro/riichi-base/src/models/constants/languages"
)

type UraDora struct {
	Count int
}

func (y UraDora) Match(p *models.Partition, c *Conditions) bool { return c.DoubleRiichi }

func (y UraDora) Han(open bool) int { return y.Count }

func (y UraDora) Name(l languages.Language) string {
	if l == languages.EN {
		return "Ura Dora"
	}
	return "裏ドラ"
}
