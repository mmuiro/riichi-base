package yakuman

import (
	"riichi-calculator/src/models"
	"riichi-calculator/src/models/constants/groups"
	"riichi-calculator/src/models/constants/suits"
	"riichi-calculator/src/models/yaku"
)

type Daisuushii struct{}

func (y Daisuushii) Match(p *models.Partition, c *yaku.Conditions) bool {
	uniqueSets := make(map[suits.Suit]bool)
	for _, mentsu := range p.Mentsu {
		if mentsu.Tiles[0].IsHonor() && (mentsu.Kind == groups.Kantsu || mentsu.Kind == groups.Koutsu) {
			uniqueSets[mentsu.Suit] = true
		}
	}
	return uniqueSets[suits.Ton] && uniqueSets[suits.Nan] && uniqueSets[suits.Xia] && uniqueSets[suits.Pei]
}

func (y Daisuushii) Value(open bool) int {
	return 2
}

func (y Daisuushii) Description() string {
	return "1 set of each wind."
}

func (y Daisuushii) Name() string {
	return "Dai Suu Shii"
}
