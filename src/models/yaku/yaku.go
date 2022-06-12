package yaku

import "riichi-calculator/src/models"

type Yaku interface {
	Match(p *models.Partition, c *Conditions) bool
	Han(open bool) int
	Description() string
}
