package yakuman

import (
	"github.com/mmuiro/riichi-base/src/models"
	"github.com/mmuiro/riichi-base/src/models/constants/groups"
	"github.com/mmuiro/riichi-base/src/models/constants/suits"
	"github.com/mmuiro/riichi-base/src/models/yaku"
)

type DaiSuushii struct{}

func (y DaiSuushii) Match(p *models.Partition, c *yaku.Conditions) bool {
	uniqueSets := make(map[suits.Suit]bool)
	for _, mentsu := range p.Mentsu {
		if mentsu.Tiles[0].IsHonor() && (mentsu.Kind == groups.Kantsu || mentsu.Kind == groups.Koutsu) {
			uniqueSets[mentsu.Suit] = true
		}
	}
	return uniqueSets[suits.Ton] && uniqueSets[suits.Nan] && uniqueSets[suits.Xia] && uniqueSets[suits.Pei]
}

func (y DaiSuushii) Value() int {
	return 2
}

func (y DaiSuushii) Description() string {
	return "1 set of each wind."
}

func (y DaiSuushii) Name() string {
	return "Dai Suu Shii"
}
