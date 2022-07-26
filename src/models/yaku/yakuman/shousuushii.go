package yakuman

import (
	"github.com/mmuiro/riichi-base/src/models"
	"github.com/mmuiro/riichi-base/src/models/constants/groups"
	"github.com/mmuiro/riichi-base/src/models/constants/languages"
	"github.com/mmuiro/riichi-base/src/models/constants/suits"
	"github.com/mmuiro/riichi-base/src/models/yaku"
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

func (y ShouSuushii) Name(l languages.Language) string {
	if l == languages.EN {
		return "Four Little Winds"
	}
	return "小四喜"
}
