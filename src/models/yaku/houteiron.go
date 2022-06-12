package yaku

import "riichi-calculator/src/models"

type Houtei struct{}

func (y *Houtei) Match(p *models.Partition, c *Conditions) bool { return c.Houtei }

func (y *Houtei) Han(open bool) int { return 1 }

func (y *Houtei) Description() string {
	return "Win by ron on the last discard."
}
