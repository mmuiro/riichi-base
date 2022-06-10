package yaku

import "riichi-calculator/src/models"

type Yaku interface {
	match(p *models.Partition, c *Conditions) bool
	han() int
	description() string
}
