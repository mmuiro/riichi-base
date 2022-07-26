package yaku

import (
	"github.com/mmuiro/riichi-base/src/models"
	"github.com/mmuiro/riichi-base/src/models/constants/languages"
)

type RinshanKaihou struct{}

func (y RinshanKaihou) Match(p *models.Partition, c *Conditions) bool { return c.Rinshan }

func (y RinshanKaihou) Han(open bool) int { return 1 }

func (y RinshanKaihou) Name(l languages.Language) string {
	if l == languages.EN {
		return "After a Kan"
	}
	return "嶺上開花"
}
