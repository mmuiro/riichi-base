package yaku

import "riichi-calculator/src/models"

type Chankan struct{}

func (y *Chankan) Match(p *models.Partition, c *Conditions) bool { return c.Chankan }

func (y *Chankan) Han(open bool) int { return 1 }

func (y *Chankan) Description() string {
	return "Win by robbing a kan."
}
