package yaku

import (
	"github.com/mmuiro/riichi-base/src/models"
	"github.com/mmuiro/riichi-base/src/models/constants/languages"
)

type JunchanTaiyao struct{}

func (y JunchanTaiyao) Match(p *models.Partition, c *Conditions) bool {
	for _, mentsu := range p.Mentsu {
		containsTerminal := false
		for _, tile := range mentsu.Tiles {
			if !tile.IsHonor() && (tile.Value == 1 || tile.Value == 9) {
				containsTerminal = true
			}
		}
		if !containsTerminal {
			return false
		}
	}
	return true
}

func (y JunchanTaiyao) Han(open bool) int {
	if open {
		return 2
	}
	return 3
}

func (y JunchanTaiyao) Name(l languages.Language) string {
	if l == languages.EN {
		return "Fully Outside Hand"
	}
	return "純全帯么"
}
