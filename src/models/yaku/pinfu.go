package yaku

import (
	"riichi-calculator/src/models"
	"riichi-calculator/src/models/constants/waits"
)

type Pinfu struct{}

func (y Pinfu) Match(p *models.Partition, c *Conditions) bool {
	return c.Menzenchin && p.Wait == waits.Ryanmen
}

func (y Pinfu) Han(open bool) int { return 1 }

func (y Pinfu) Description() string {
	return "Win on a ryanmen, with no yakuhai/fanpai tiles in hand."
}

func (y Pinfu) Name() string {
	return "Pinfu"
}
