package yaku

import (
	"github.com/mmuiro/riichi-base/src/models"
	"github.com/mmuiro/riichi-base/src/models/constants/languages"
)

type Chankan struct{}

func (y Chankan) Match(p *models.Partition, c *Conditions) bool { return c.Chankan }

func (y Chankan) Han(open bool) int { return 1 }

func (y Chankan) Name(l languages.Language) string {
	if l == languages.EN {
		return "Robbing a Kan"
	}
	return "槍槓"
}
