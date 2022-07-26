package yakuman

import (
	"github.com/mmuiro/riichi-base/src/models"
	"github.com/mmuiro/riichi-base/src/models/constants/languages"
	"github.com/mmuiro/riichi-base/src/models/constants/waits"
	"github.com/mmuiro/riichi-base/src/models/yaku"
)

type KokushiThirteen struct{}

func (y KokushiThirteen) Match(p *models.Partition, c *yaku.Conditions) bool {
	return models.CheckKokushi(p) && p.Wait == waits.KokushiThirteen
}

func (y KokushiThirteen) Value() int {
	return 2
}

func (y KokushiThirteen) Name(l languages.Language) string {
	if l == languages.EN {
		return "Thirteen-wait Thirteen Orphans"
	}
	return "国士無双13面待"
}
