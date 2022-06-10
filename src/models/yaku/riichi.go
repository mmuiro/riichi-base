package yaku

import "riichi-calculator/src/models"

type Riichi struct{}

func (y *Riichi) match(p *models.Partition, c *Conditions) bool { return c.Riichi }

func (y *Riichi) han() int { return 1 }

func (y *Riichi) description() string {
	return "Win after calling riichi from a closed hand on tenpai."
}
