package yakuman

import (
	"github.com/mmuiro/riichi-base/src/models"
	"github.com/mmuiro/riichi-base/src/models/constants/languages"
	"github.com/mmuiro/riichi-base/src/models/constants/waits"
	"github.com/mmuiro/riichi-base/src/models/yaku"
)

type JunseiChuuren struct{}

func (y JunseiChuuren) Match(p *models.Partition, c *yaku.Conditions) bool {
	if !c.Menzenchin {
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
	return p.Wait == waits.JunseiChuuren

}

func (y JunseiChuuren) Value() int {
	return 2
}

func (y JunseiChuuren) Name(l languages.Language) string {
	if l == languages.EN {
		return "Pure Nine Gates"
	}
	return "純正九蓮宝燈"
}
