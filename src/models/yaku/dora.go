package yaku

import (
	"github.com/mmuiro/riichi-base/src/models"
	"github.com/mmuiro/riichi-base/src/models/constants/languages"
)

type Dora struct {
	Count int
}

func (y Dora) Match(p *models.Partition, c *Conditions) bool { return c.DoubleRiichi }

func (y Dora) Han(open bool) int { return y.Count }

func (y Dora) Name(l languages.Language) string {
	if l == languages.EN {
		return "Dora"
	}
	return "ドラ"
}
