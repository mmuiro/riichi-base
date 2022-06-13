package yaku

import "riichi-calculator/src/models"

type DoubleRiichi struct{}

func (y DoubleRiichi) Match(p *models.Partition, c *Conditions) bool { return c.DoubleRiichi }

func (y DoubleRiichi) Han(open bool) int { return 2 }

func (y DoubleRiichi) Description() string {
	return "Win after calling riichi on your first discard from a closed hand on tenpai."
}

func (y DoubleRiichi) Name() string {
	return "Double Riichi"
}
