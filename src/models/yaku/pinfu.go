package yaku

import (
	"github.com/mmuiro/riichi-base/src/models"
	"github.com/mmuiro/riichi-base/src/models/constants/groups"
	"github.com/mmuiro/riichi-base/src/models/constants/suits"
	"github.com/mmuiro/riichi-base/src/models/constants/waits"
)

type Pinfu struct{}

func (y Pinfu) Match(p *models.Partition, c *Conditions) bool {
	models.CheckAndAssignMentsuCounts(p)
	for _, m := range p.Mentsu {
		if m.Kind == groups.Toitsu && (m.Suit == c.Jikaze || m.Suit == c.Bakaze ||
			m.Suit == suits.Chun || m.Suit == suits.Haku || m.Suit == suits.Hatsu) {
			return false
		}
	}
	return c.Menzenchin && p.Wait == waits.Ryanmen && p.MentsuCounts[groups.Shuntsu] == 4
}

func (y Pinfu) Han(open bool) int { return 1 }

func (y Pinfu) Description() string {
	return "Win on a ryanmen, with no yakuhai/fanpai tiles in hand."
}

func (y Pinfu) Name() string {
	return "Pinfu"
}
