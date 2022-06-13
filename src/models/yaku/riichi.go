package yaku

import "riichi-calculator/src/models"

type Riichi struct{}

func (y Riichi) Match(p *models.Partition, c *Conditions) bool { return c.Riichi }

func (y Riichi) Han(open bool) int { return 1 }

func (y Riichi) Description() string {
	return "Win after calling riichi from a closed hand on tenpai."
}

func (y Riichi) Name() string {
	return "Riichi"
}
