package yakuman

import (
	"github.com/mmuiro/riichi-base/src/models"
	"github.com/mmuiro/riichi-base/src/models/constants/languages"
	"github.com/mmuiro/riichi-base/src/models/yaku"
	"github.com/mmuiro/riichi-base/src/utils"
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

func (y TsuuIisou) Name(l languages.Language) string {
	if l == languages.EN {
		return "All Honors"
	}
	return "字一色"
}
