package yakuman

import (
	"riichi-calculator/src/models"
	"riichi-calculator/src/models/constants/groups"
	"riichi-calculator/src/models/constants/waits"
	"riichi-calculator/src/models/yaku"
)

type SuuAnkouTanki struct{}

func (y SuuAnkouTanki) Match(p *models.Partition, c *yaku.Conditions) bool {
	concealedSets := 0
	for _, mentsu := range p.Mentsu {
		if (mentsu.Kind == groups.Kantsu || mentsu.Kind == groups.Koutsu) && !mentsu.Open {
			concealedSets++
		}
	}
	return concealedSets == 4 && p.Wait == waits.Tanki
}

func (y SuuAnkouTanki) Value() int {
	return 2
}

func (y SuuAnkouTanki) Description() string {
	return "4 concealed sets, single wait."
}

func (y SuuAnkouTanki) Name() string {
	return "Suu Ankou Tanki"
}
