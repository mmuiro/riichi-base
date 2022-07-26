package yaku

import (
	"github.com/mmuiro/riichi-base/src/models"
	"github.com/mmuiro/riichi-base/src/models/constants/languages"
)

type Houtei struct{}

func (y Houtei) Match(p *models.Partition, c *Conditions) bool { return c.Houtei }

func (y Houtei) Han(open bool) int { return 1 }

func (y Houtei) Name(l languages.Language) string {
	if l == languages.EN {
		return "Under the River"
	}
	return "河底撈魚"
}
