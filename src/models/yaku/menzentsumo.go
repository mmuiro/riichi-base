package yaku

import "github.com/mmuiro/riichi-base/src/models"

type MenzenTsumo struct{}

func (y MenzenTsumo) Match(p *models.Partition, c *Conditions) bool {
	return c.Menzenchin && c.Tsumo
}

func (y MenzenTsumo) Han(open bool) int { return 1 }

func (y MenzenTsumo) Description() string {
	return "Tsumo on a closed hand."
}

func (y MenzenTsumo) Name() string {
	return "Menzenchin Tsumo"
}
