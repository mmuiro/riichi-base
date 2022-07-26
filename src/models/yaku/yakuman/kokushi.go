package yakuman

import (
	"github.com/mmuiro/riichi-base/src/models"
	"github.com/mmuiro/riichi-base/src/models/constants/languages"
	"github.com/mmuiro/riichi-base/src/models/constants/waits"
	"github.com/mmuiro/riichi-base/src/models/yaku"
)

type Kokushi struct{}

func (y Kokushi) Match(p *models.Partition, c *yaku.Conditions) bool {
	return models.CheckKokushi(p) && p.Wait == waits.KokushiSingle
}

func (y Kokushi) Value() int {
	return 1
}

func (y Kokushi) Name(l languages.Language) string {
	if l == languages.EN {
		return "Thirteen Orphans"
	}
	return "国士無双"
}
