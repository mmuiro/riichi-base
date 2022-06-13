package yakuman

import (
	"riichi-calculator/src/models"
	"riichi-calculator/src/models/constants/waits"
	"riichi-calculator/src/models/yaku"
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
	return 1
}

func (y JunseiChuuren) Description() string {
	return "Nine gates, 1112345678999+X pattern of one suit, 1-9 wait."
}

func (y JunseiChuuren) Name() string {
	return "Junsei Chuuren Poutou"
}
