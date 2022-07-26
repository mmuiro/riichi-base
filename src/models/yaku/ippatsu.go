package yaku

import (
	"github.com/mmuiro/riichi-base/src/models"
	"github.com/mmuiro/riichi-base/src/models/constants/languages"
)

type Ippatsu struct{}

func (y Ippatsu) Match(p *models.Partition, c *Conditions) bool { return c.Ippatsu }

func (y Ippatsu) Han(open bool) int { return 1 }

func (y Ippatsu) Name(l languages.Language) string {
	if l == languages.EN {
		return "Ippatsu"
	}
	return "一発"
}
