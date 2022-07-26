package yaku

import (
	"github.com/mmuiro/riichi-base/src/models"
	"github.com/mmuiro/riichi-base/src/models/constants/groups"
	"github.com/mmuiro/riichi-base/src/models/constants/languages"
)

type SanKantsu struct{}

func (y SanKantsu) Match(p *models.Partition, c *Conditions) bool {
	models.CheckAndAssignMentsuCounts(p)
	return p.MentsuCounts[groups.Kantsu] == 3
}

func (y SanKantsu) Han(open bool) int {
	return 2
}

func (y SanKantsu) Name(l languages.Language) string {
	if l == languages.EN {
		return "Three Quads"
	}
	return "三槓子"
}
