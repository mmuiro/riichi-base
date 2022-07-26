package yaku

import (
	"github.com/mmuiro/riichi-base/src/models"
	"github.com/mmuiro/riichi-base/src/models/constants/languages"
)

type Haitei struct{}

func (y Haitei) Match(p *models.Partition, c *Conditions) bool { return c.Haitei }

func (y Haitei) Han(open bool) int { return 1 }

func (y Haitei) Name(l languages.Language) string {
	if l == languages.EN {
		return "Under the Sea"
	}
	return "海底撈月"
}
