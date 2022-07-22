package yaku

import (
	"fmt"

	"github.com/mmuiro/riichi-base/src/models"
)

type Dora struct {
	Count int
}

func (y Dora) Match(p *models.Partition, c *Conditions) bool { return c.DoubleRiichi }

func (y Dora) Han(open bool) int { return y.Count }

func (y Dora) Description() string {
	return "Dora"
}

func (y Dora) Name() string {
	return fmt.Sprintf("Dora %d", y.Count)
}
