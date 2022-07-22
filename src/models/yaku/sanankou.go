package yaku

import (
	"github.com/mmuiro/riichi-base/src/models"
	"github.com/mmuiro/riichi-base/src/models/constants/groups"
)

type SanAnkou struct{}

func (y SanAnkou) Match(p *models.Partition, c *Conditions) bool {
	closedSets := 0
	for _, mentsu := range p.Mentsu {
		if mentsu.Kind == groups.Kantsu || mentsu.Kind == groups.Koutsu && !mentsu.Open {
			closedSets++
		}
	}
	return closedSets == 3
}

func (y SanAnkou) Han(open bool) int {
	return 2
}

func (y SanAnkou) Description() string {
	return "Three concealed triplets."
}

func (y SanAnkou) Name() string {
	return "San Ankou"
}
