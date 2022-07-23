package yaku

import (
	"github.com/mmuiro/riichi-base/src/models"
)

type AkaDora struct {
	Count int
}

func (y AkaDora) Match(p *models.Partition, c *Conditions) bool { return c.DoubleRiichi }

func (y AkaDora) Han(open bool) int { return y.Count }

func (y AkaDora) Description() string {
	return "Red Dora"
}

func (y AkaDora) Name() string {
	return "Aka Dora"
}
