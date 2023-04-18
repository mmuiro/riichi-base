package yaku

import (
	"github.com/mmuiro/riichi-base/src/models"
	"github.com/mmuiro/riichi-base/src/models/constants/languages"
)

type Chanta struct{}

func (y Chanta) Match(p *models.Partition, c *Conditions) bool {
    hasHonor := false
	for _, mentsu := range p.Mentsu {
		check := false
		for _, tile := range mentsu.Tiles {
			if tile.IsHonor() || tile.Value == 1 || tile.Value == 9 {
                if tile.IsHonor() {
                    hasHonor = true
                }
				check = true
			}
		}
		if !check {
			return false
		}
	}
	return hasHonor
}

func (y Chanta) Han(open bool) int {
	if open {
		return 1
	}
	return 2
}

func (y Chanta) Name(l languages.Language) string {
	if l == languages.EN {
		return "Half Outside Hand"
	}
	return "全帯幺九"
}
