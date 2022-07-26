package yakuman

import (
	"github.com/mmuiro/riichi-base/src/models"
	"github.com/mmuiro/riichi-base/src/models/constants/languages"
	"github.com/mmuiro/riichi-base/src/models/yaku"
)

type Tenhou struct{}

func (y Tenhou) Match(p *models.Partition, c *yaku.Conditions) bool {
	return c.Tenhou
}

func (y Tenhou) Value() int {
	return 1
}

func (y Tenhou) Name(l languages.Language) string {
	if l == languages.EN {
		return "Heaven's Blessing"
	}
	return "天和"
}
