package yakuman

import (
	"riichi-calculator/src/models"
	"riichi-calculator/src/models/constants/groups"
	"riichi-calculator/src/models/constants/suits"
	"riichi-calculator/src/models/yaku"
)

type ShouSuushii struct{}

func (y ShouSuushii) Match(p *models.Partition, c *yaku.Conditions) bool {
	uniquePairs, uniqueSets := make(map[suits.Suit]bool), make(map[suits.Suit]bool)
	for _, mentsu := range p.Mentsu {
		if mentsu.Tiles[0].IsHonor() {
			if mentsu.Kind == groups.Kantsu || mentsu.Kind == groups.Koutsu {
				uniqueSets[mentsu.Suit] = true
			} else if mentsu.Kind == groups.Toitsu {
				uniquePairs[mentsu.Suit] = true
			}
		}
	}
	return (uniqueSets[suits.Ton] && uniqueSets[suits.Nan] && uniqueSets[suits.Xia] && uniquePairs[suits.Pei]) ||
		(uniqueSets[suits.Pei] && uniqueSets[suits.Nan] && uniqueSets[suits.Xia] && uniquePairs[suits.Ton]) ||
		(uniqueSets[suits.Ton] && uniqueSets[suits.Pei] && uniqueSets[suits.Xia] && uniquePairs[suits.Nan]) ||
		(uniqueSets[suits.Ton] && uniqueSets[suits.Nan] && uniqueSets[suits.Pei] && uniquePairs[suits.Xia])
}

func (y ShouSuushii) Value() int {
	return 1
}

func (y ShouSuushii) Description() string {
	return "3 sets and 1 pair of winds."
}

func (y ShouSuushii) Name() string {
	return "Shou Suu Shii"
}
