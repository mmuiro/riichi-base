package yaku

import (
	"github.com/mmuiro/riichi-base/src/models"
	"github.com/mmuiro/riichi-base/src/utils"
)

type HonRoutou struct{}

func (y HonRoutou) Match(p *models.Partition, c *Conditions) bool {
	tiles := p.Tiles()
	return utils.All(utils.FuncMap(func(t models.Tile) bool {
		return t.IsHonor() || t.Value == 1 || t.Value == 9
	}, tiles))
}

func (y HonRoutou) Han(open bool) int { return 2 }

func (y HonRoutou) Description() string {
	return "All terminals/honors."
}

func (y HonRoutou) Name() string {
	return "Honroutou"
}
