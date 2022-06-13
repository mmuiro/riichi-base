package yaku

import "riichi-calculator/src/models"

type Haitei struct{}

func (y Haitei) Match(p *models.Partition, c *Conditions) bool { return c.Haitei }

func (y Haitei) Han(open bool) int { return 1 }

func (y Haitei) Description() string {
	return "Win on the last draw from the wall."
}

func (y Haitei) Name() string {
	return "Win on the last draw from the wall."
}
