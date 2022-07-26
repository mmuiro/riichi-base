package yaku

import (
	"github.com/mmuiro/riichi-base/src/models"
	"github.com/mmuiro/riichi-base/src/models/constants/groups"
	"github.com/mmuiro/riichi-base/src/models/constants/languages"
	"github.com/mmuiro/riichi-base/src/models/constants/suits"
)

type ShouSangen struct{}

func (y ShouSangen) Match(p *models.Partition, c *Conditions) bool {
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
	return (uniqueSets[suits.Chun] && uniqueSets[suits.Haku] && uniquePairs[suits.Hatsu]) ||
		(uniqueSets[suits.Hatsu] && uniqueSets[suits.Haku] && uniquePairs[suits.Chun]) ||
		(uniqueSets[suits.Chun] && uniqueSets[suits.Hatsu] && uniquePairs[suits.Haku])
}

func (y ShouSangen) Han(open bool) int {
	return 2
}

func (y ShouSangen) Name(l languages.Language) string {
	if l == languages.EN {
		return "Little Three Dragons"
	}
	return "小三元"
}
