package yakuman

import (
	"riichi-calculator/src/models"
	"riichi-calculator/src/models/yaku"
	"riichi-calculator/src/utils"
)

type TsuuIisou struct{}

func (y TsuuIisou) Match(p *models.Partition, c *yaku.Conditions) bool {
	return utils.All(utils.FuncMap(func(t models.Tile) bool {
		return t.IsHonor()
	}, p.Tiles()))
}

func (y TsuuIisou) Value() int {
	return 1
}

func (y TsuuIisou) Description() string {
	return "All honors."
}

func (y TsuuIisou) Name() string {
	return "Tsuuiisou"
}
