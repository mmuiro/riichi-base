package yakuman

import (
	"riichi-calculator/src/models"
	"riichi-calculator/src/models/yaku"
	"riichi-calculator/src/utils"
)

type Chinroutou struct{}

func (y Chinroutou) Match(p *models.Partition, c *yaku.Conditions) bool {
	return utils.All(utils.FuncMap(func(t models.Tile) bool {
		return !t.IsHonor() && t.Value == 1 || t.Value == 9
	}, p.Tiles()))
}

func (y Chinroutou) Value() int {
	return 1
}

func (y Chinroutou) Description() string {
	return "All terminals"
}

func (y Chinroutou) Name() string {
	return "Chinroutou"
}
