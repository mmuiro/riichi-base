package yakuman

import (
	"github.com/mmuiro/riichi-base/src/models"
	"github.com/mmuiro/riichi-base/src/models/constants/groups"
	"github.com/mmuiro/riichi-base/src/models/constants/languages"
	"github.com/mmuiro/riichi-base/src/models/yaku"
)

type SuuKantsu struct{}

func (y SuuKantsu) Match(p *models.Partition, c *yaku.Conditions) bool {
	models.CheckAndAssignMentsuCounts(p)
	return p.MentsuCounts[groups.Kantsu] == 4
}

func (y SuuKantsu) Value() int {
	return 1
}

func (y SuuKantsu) Name(l languages.Language) string {
	if l == languages.EN {
		return "Four Quads"
	}
	return "四槓子"
}
