package yakuman

import (
	"github.com/mmuiro/riichi-base/src/models"
	"github.com/mmuiro/riichi-base/src/models/constants/languages"
	"github.com/mmuiro/riichi-base/src/models/yaku"
	"github.com/mmuiro/riichi-base/src/utils"
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

func (y Chinroutou) Name(l languages.Language) string {
	if l == languages.EN {
		return "All Terminals"
	}
	return "清老頭"
}
