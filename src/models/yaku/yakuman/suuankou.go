package yakuman

import (
	"riichi-calculator/src/models"
	"riichi-calculator/src/models/constants/groups"
	"riichi-calculator/src/models/constants/waits"
	"riichi-calculator/src/models/yaku"
)

type SuuAnkou struct{}

func (y SuuAnkou) Match(p *models.Partition, c *yaku.Conditions) bool {
	concealedSets := 0
	for _, mentsu := range p.Mentsu {
		if (mentsu.Kind == groups.Kantsu || mentsu.Kind == groups.Koutsu) && !mentsu.Open {
			concealedSets++
		}
	}
	return concealedSets == 4 && p.Wait == waits.Shanpon
}

func (y SuuAnkou) Value() int {
	return 1
}

func (y SuuAnkou) Description() string {
	return "4 concealed sets."
}

func (y SuuAnkou) Name() string {
	return "Suu Ankou"
}
