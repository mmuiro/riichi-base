package yaku

import (
	"github.com/mmuiro/riichi-base/src/models"
	"github.com/mmuiro/riichi-base/src/models/constants/languages"
)

type AkaDora struct {
	Count int
}

func (y AkaDora) Match(p *models.Partition, c *Conditions) bool { return c.DoubleRiichi }

func (y AkaDora) Han(open bool) int { return y.Count }

func (y AkaDora) Name(l languages.Language) string {
	if l == languages.EN {
		return "Red Five"
	}
	return "赤ドラ"
}
