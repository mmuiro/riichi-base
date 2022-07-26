package yaku

import (
	"github.com/mmuiro/riichi-base/src/models"
	"github.com/mmuiro/riichi-base/src/models/constants/languages"
)

type MenzenTsumo struct{}

func (y MenzenTsumo) Match(p *models.Partition, c *Conditions) bool {
	return c.Menzenchin && c.Tsumo
}

func (y MenzenTsumo) Han(open bool) int { return 1 }

func (y MenzenTsumo) Name(l languages.Language) string {
	if l == languages.EN {
		return "Fully Concealed Hand"
	}
	return "門前清自摸和"
}
