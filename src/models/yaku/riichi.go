package yaku

import (
	"github.com/mmuiro/riichi-base/src/models"
	"github.com/mmuiro/riichi-base/src/models/constants/languages"
)

type Riichi struct{}

func (y Riichi) Match(p *models.Partition, c *Conditions) bool { return c.Riichi }

func (y Riichi) Han(open bool) int { return 1 }

func (y Riichi) Name(l languages.Language) string {
	if l == languages.EN {
		return "Riichi"
	}
	return "立直"
}
