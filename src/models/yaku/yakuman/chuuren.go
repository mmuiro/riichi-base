package yakuman

import (
	"riichi-calculator/src/models"
	"riichi-calculator/src/models/constants/waits"
	"riichi-calculator/src/models/yaku"
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

func (y Chuuren) Value(open bool) int {
	return 1
}

func (y Chuuren) Description() string {
	return "Nine gates, 1112345678999+X pattern of one suit."
}

func (y Chuuren) Name() string {
	return "Chuuren Poutou"
}
