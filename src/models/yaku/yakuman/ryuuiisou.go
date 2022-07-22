package yakuman

import (
	"github.com/mmuiro/riichi-base/src/models"
	"github.com/mmuiro/riichi-base/src/models/yaku"
	"github.com/mmuiro/riichi-base/src/utils"
)

type RyuuIisou struct{}

func (y RyuuIisou) Match(p *models.Partition, c *yaku.Conditions) bool {
	possibleIDs := []int{10, 11, 12, 14, 16, 32}
	return utils.All(utils.FuncMap(func(t models.Tile) bool {
		return utils.Contains(possibleIDs, models.TileToID(&t))
	}, p.Tiles()))
}

func (y RyuuIisou) Value() int {
	return 1
}

func (y RyuuIisou) Description() string {
	return "All green tiles."
}

func (y RyuuIisou) Name() string {
	return "Ryuuiisou"
}
