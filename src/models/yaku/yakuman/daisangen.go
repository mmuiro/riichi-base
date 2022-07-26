package yakuman

import (
	"github.com/mmuiro/riichi-base/src/models"
	"github.com/mmuiro/riichi-base/src/models/constants/groups"
	"github.com/mmuiro/riichi-base/src/models/constants/languages"
	"github.com/mmuiro/riichi-base/src/models/constants/suits"
	"github.com/mmuiro/riichi-base/src/models/yaku"
)

type DaiSangen struct{}

func (y DaiSangen) Match(p *models.Partition, c *yaku.Conditions) bool {
	uniqueSets := make(map[suits.Suit]bool)
	for _, mentsu := range p.Mentsu {
		if mentsu.Tiles[0].IsHonor() && (mentsu.Kind == groups.Kantsu || mentsu.Kind == groups.Koutsu) {
			uniqueSets[mentsu.Suit] = true
		}
	}
	return uniqueSets[suits.Chun] && uniqueSets[suits.Haku] && uniqueSets[suits.Hatsu]
}

func (y DaiSangen) Value() int {
	return 1
}

func (y DaiSangen) Name(l languages.Language) string {
	if l == languages.EN {
		return "Big Three Dragons"
	}
	return "大三元"
}
