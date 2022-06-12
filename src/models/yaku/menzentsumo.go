package yaku

import "riichi-calculator/src/models"

type MenzenTsumo struct{}

func (y *MenzenTsumo) Match(p *models.Partition, c *Conditions) bool {
	return c.Menzenchin && c.Tsumo
}

func (y *MenzenTsumo) Han(open bool) int { return 1 }

func (y *MenzenTsumo) Description() string {
	return "Tsumo on a closed hand."
}
