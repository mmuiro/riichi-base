package yakuman

import (
	"riichi-calculator/src/models"
	"riichi-calculator/src/models/constants/groups"
	"riichi-calculator/src/models/constants/suits"
	"riichi-calculator/src/models/yaku"
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

func (y DaiSangen) Value(open bool) int {
	return 1
}

func (y DaiSangen) Description() string {
	return "1 set of each dragon."
}

func (y DaiSangen) Name() string {
	return "Dai Sangen"
}
