package yaku

import (
	"github.com/mmuiro/riichi-base/src/models"
	"github.com/mmuiro/riichi-base/src/models/constants/groups"
)

type Toitoi struct{}

func (y Toitoi) Match(p *models.Partition, c *Conditions) bool {
	models.CheckAndAssignMentsuCounts(p)
	return len(p.Mentsu) == 5 &&
		p.MentsuCounts[groups.Kantsu]+p.MentsuCounts[groups.Koutsu] == 4 &&
		p.MentsuCounts[groups.Toitsu] == 1
}

func (y Toitoi) Han(open bool) int {
	return 2
}

func (y Toitoi) Description() string {
	return "All sets and a pair."
}

func (y Toitoi) Name() string {
	return "Toitoi"
}
