package yakuman

import (
	"riichi-calculator/src/models"
	"riichi-calculator/src/models/constants/groups"
	"riichi-calculator/src/models/yaku"
)

type SuuKantsu struct{}

func (y SuuKantsu) Match(p *models.Partition, c *yaku.Conditions) bool {
	models.CheckAndAssignMentsuCounts(p)
	return p.MentsuCounts[groups.Kantsu] == 4
}

func (y SuuKantsu) Value(open bool) int {
	return 1
}

func (y SuuKantsu) Description() string {
	return "4 kans."
}

func (y SuuKantsu) Name() string {
	return "Suu Kantsu"
}
