package yaku

import "github.com/mmuiro/riichi-base/src/models"

type Ippatsu struct{}

func (y Ippatsu) Match(p *models.Partition, c *Conditions) bool { return c.Ippatsu }

func (y Ippatsu) Han(open bool) int { return 1 }

func (y Ippatsu) Description() string {
	return "Win on or before your next draw after calling riichi, without others calling tiles."
}

func (y Ippatsu) Name() string {
	return "Ippatsu"
}
