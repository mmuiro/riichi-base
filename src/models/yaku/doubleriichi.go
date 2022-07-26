package yaku

import (
	"github.com/mmuiro/riichi-base/src/models"
	"github.com/mmuiro/riichi-base/src/models/constants/languages"
)

type DoubleRiichi struct{}

func (y DoubleRiichi) Match(p *models.Partition, c *Conditions) bool { return c.DoubleRiichi }

func (y DoubleRiichi) Han(open bool) int { return 2 }

func (y DoubleRiichi) Name(l languages.Language) string {
	if l == languages.EN {
		return "Double Riichi"
	}
	return "ダブル立直"
}
