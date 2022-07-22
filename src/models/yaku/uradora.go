package yaku

import (
	"fmt"

	"github.com/mmuiro/riichi-base/src/models"
)

type UraDora struct {
	Count int
}

func (y UraDora) Match(p *models.Partition, c *Conditions) bool { return c.DoubleRiichi }

func (y UraDora) Han(open bool) int { return y.Count }

func (y UraDora) Description() string {
	return "Ura Dora"
}

func (y UraDora) Name() string {
	return fmt.Sprintf("Ura Dora %d", y.Count)
}
