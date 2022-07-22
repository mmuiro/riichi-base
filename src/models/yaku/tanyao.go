package yaku

import (
	"github.com/mmuiro/riichi-base/src/models"
	"github.com/mmuiro/riichi-base/src/utils"
)

type Tanyao struct{}

func (y Tanyao) Match(p *models.Partition, c *Conditions) bool {
	tiles := p.Tiles()
	return utils.All(utils.FuncMap(func(t models.Tile) bool {
		return !t.IsHonor() && 1 < t.Value && t.Value < 9
	}, tiles))
}

func (y Tanyao) Han(open bool) int { return 1 }

func (y Tanyao) Description() string { return "All simples." }

func (y Tanyao) Name() string { return "Tanyao" }
