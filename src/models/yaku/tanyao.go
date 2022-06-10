package yaku

import (
	"riichi-calculator/src/models"
	"riichi-calculator/src/utils"
)

type Tanyao struct{}

func (y *Tanyao) match(p *models.Partition, c *Conditions) bool {
	tiles := p.Tiles()
	return utils.All(utils.FuncMap(func(t models.Tile) bool {
		return !t.IsHonor() && 1 < t.Value && t.Value < 9
	}, tiles))
}

func (y *Tanyao) han() int { return 1 }

func (y *Tanyao) description() string { return "All simples." }
