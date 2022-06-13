package yaku

import "riichi-calculator/src/models"

type RinshanKaihou struct{}

func (y RinshanKaihou) Match(p *models.Partition, c *Conditions) bool { return c.Rinshan }

func (y RinshanKaihou) Han(open bool) int { return 1 }

func (y RinshanKaihou) Description() string {
	return "Win on tsumo directly after calling kan."
}

func (y RinshanKaihou) Name() string {
	return "Rinshan Kaihou"
}
