package yakuman

import (
	"github.com/mmuiro/riichi-base/src/models"
	"github.com/mmuiro/riichi-base/src/models/constants/languages"
	"github.com/mmuiro/riichi-base/src/models/yaku"
)

type Chiihou struct{}

func (y Chiihou) Match(p *models.Partition, c *yaku.Conditions) bool {
	return c.Chiihou
}

func (y Chiihou) Value() int {
	return 1
}

func (y Chiihou) Name(l languages.Language) string {
	if l == languages.EN {
		return "Earth's Blessing"
	}
	return "地和"
}
