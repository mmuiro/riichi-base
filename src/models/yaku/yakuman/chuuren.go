package yakuman

import (
	"github.com/mmuiro/riichi-base/src/models"
	"github.com/mmuiro/riichi-base/src/models/constants/languages"
	"github.com/mmuiro/riichi-base/src/models/constants/waits"
	"github.com/mmuiro/riichi-base/src/models/yaku"
)

type Chuuren struct{}

func (y Chuuren) Match(p *models.Partition, c *yaku.Conditions) bool {
	if p.Wait == waits.JunseiChuuren || !c.Menzenchin {
		return false
	}
	suit := p.Mentsu[0].Suit
	valueCounts := make([]int, 9)
	for _, tile := range p.Tiles() {
		if tile.Suit == suit && !tile.IsHonor() {
			valueCounts[tile.Value-1]++
		}
	}
	thresholds := []int{3, 1, 1, 1, 1, 1, 1, 1, 3}
	for i := 0; i < 9; i++ {
		if valueCounts[i] < thresholds[i] {
			return false
		}
	}
	return true

}

func (y Chuuren) Value() int {
	return 1
}

func (y Chuuren) Name(l languages.Language) string {
	if l == languages.EN {
		return "Nine Gates"
	}
	return "九連宝燈"
}
