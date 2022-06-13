package yaku

import (
	"riichi-calculator/src/models"
	"riichi-calculator/src/models/constants/groups"
)

type SanKantsu struct{}

func (y SanKantsu) Match(p *models.Partition, c *Conditions) bool {
	models.CheckAndAssignMentsuCounts(p)
	return p.MentsuCounts[groups.Kantsu] == 3
}

func (y SanKantsu) Han(open bool) int {
	return 2
}

func (y SanKantsu) Description() string {
	return "Three kans."
}

func (y SanKantsu) Name() string {
	return "San Kantsu"
}
