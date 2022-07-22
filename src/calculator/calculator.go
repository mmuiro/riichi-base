package calculator

import (
	"fmt"
	"math"

	"github.com/mmuiro/riichi-base/src/models"
	"github.com/mmuiro/riichi-base/src/models/constants/groups"
	"github.com/mmuiro/riichi-base/src/models/constants/suits"
	"github.com/mmuiro/riichi-base/src/models/constants/waits"
	"github.com/mmuiro/riichi-base/src/models/yaku"
	"github.com/mmuiro/riichi-base/src/models/yaku/yakuman"
)

type NoAgariError struct {
	h *models.Hand
}

func (e *NoAgariError) Error() string {
	return fmt.Sprintf("The hand %s is incomplete.", e.h.String())
}

type NoYakuError struct {
	h *models.Hand
}

func (e *NoYakuError) Error() string {
	return fmt.Sprintf("The hand %s has no yaku.", e.h.String())
}

func roundUp(val int, inc int) int {
	r := val % inc
	if r > 0 {
		return val + inc - r
	}
	return val
}

// Returns the calculated fu for the given partition and conditions.
func CalculateFu(p *models.Partition, c *yaku.Conditions) int {
	if models.CheckChiiToitsu(p) {
		return 25
	}
	pinfu := yaku.Pinfu{}.Match(p, c)
	fu := 20
	for _, mentsu := range p.Mentsu {
		mentsuFu := 0
		if mentsu.Kind == groups.Koutsu {
			mentsuFu = 2
		} else if mentsu.Kind == groups.Kantsu {
			mentsuFu = 8
		}
		if !mentsu.Open {
			mentsuFu *= 2
		}
		if mentsu.Tiles[0].IsHonor() || mentsu.Tiles[0].Value == 1 || mentsu.Tiles[0].Value == 9 {
			mentsuFu *= 2
		}
		// yakuhai pair check
		if mentsu.Kind == groups.Toitsu {
			if mentsu.Suit == suits.Chun || mentsu.Suit == suits.Haku || mentsu.Suit == suits.Hatsu || mentsu.Suit == c.Jikaze {
				mentsuFu = 2
			}
			if mentsu.Suit == c.Bakaze {
				mentsuFu += 2
			}
		}
		fu += mentsuFu
	}
	if p.Wait == waits.Kanchan || p.Wait == waits.Penchan || p.Wait == waits.Tanki {
		fu += 2
	}
	// clean this up
	if c.Tsumo && !pinfu {
		fu += 2
	} else if c.Menzenchin && !c.Tsumo {
		fu += 10
	}
	if fu == 20 && !pinfu { // open hand ron with no bonus fu forced to 30 fu
		fu += 10
	}
	fu = roundUp(fu, 10)
	return fu
}

func FindHanAndYaku(p *models.Partition, c *yaku.Conditions) (int, []yaku.Yaku) {
	han, yakuList := 0, make([]yaku.Yaku, 0)
	for _, y := range yaku.AllYaku {
		if y.Match(p, c) {
			han += y.Han(!c.Menzenchin)
			yakuList = append(yakuList, y)
		}
	}
	return han, yakuList
}

func FindYakuman(p *models.Partition, c *yaku.Conditions) (int, []yakuman.Yakuman) {
	multiplier, yakumanList := 0, make([]yakuman.Yakuman, 0)
	for _, y := range yakuman.AllYakuman {
		if y.Match(p, c) {
			multiplier += y.Value()
			yakumanList = append(yakumanList, y)
		}
	}
	return multiplier, yakumanList
}

func CalculateScore(h *models.Hand, w *models.Tile, c *yaku.Conditions) (int, error) {
	score, _, _, _, _, _, _, err := CalculateScoreVerbose(h, w, c)
	if err != nil {
		return 0, err
	}
	return score, nil
}

func CalculateScoreVerbose(h *models.Hand, w *models.Tile, c *yaku.Conditions) (int, *models.Partition, []yaku.Yaku, []yakuman.Yakuman, int, int, string, error) {
	agari, partitions := models.CheckAgari(h, w, c.Tsumo)
	if !agari {
		return 0, nil, nil, nil, 0, 0, "", &NoAgariError{h: h}
	}
	maxScore := 0
	var maxPartition models.Partition
	var maxYakuList []yaku.Yaku
	var maxYakumanList []yakuman.Yakuman
	var maxSLevel string
	var maxHan, maxFu int
	for _, p := range partitions {
		score, yakuList, yakumanList, han, fu, slevel := calculatePartitionScore(&p, c)
		if score > maxScore {
			maxScore, maxPartition, maxYakuList, maxYakumanList, maxSLevel = score, p, yakuList, yakumanList, slevel
			maxHan, maxFu = han, fu
		}
	}
	if maxScore == 0 {
		return 0, nil, nil, nil, 0, 0, "", &NoYakuError{h: h}
	}
	return maxScore, &maxPartition, maxYakuList, maxYakumanList, maxHan, maxFu, maxSLevel, nil
}

func calculatePartitionScore(p *models.Partition, c *yaku.Conditions) (int, []yaku.Yaku, []yakuman.Yakuman, int, int, string) {
	yakumanMultiplier, yakumanList := FindYakuman(p, c)
	score := 0
	// has a yakuman
	if yakumanMultiplier > 0 {
		if c.Jikaze == suits.Ton {
			score = 6 * ScoreLevelToBasicPoints[Yakuman]
		} else {
			score = 4 * ScoreLevelToBasicPoints[Yakuman]
		}
		return yakumanMultiplier * score, nil, yakumanList, 0, 0, "Yakuman"
	} else {
		// Find the han
		han, yakuList := FindHanAndYaku(p, c)
		var fu int
		if han == 0 {
			return 0, nil, nil, 0, 0, ""
		}
		// if the hand has yaku, add dora and red dora, and uradora if menzenchin
		dora, akadora, uradora := 0, 0, 0
		for _, tile := range p.Tiles() {
			if tile.Red {
				akadora++
			}
			for _, d := range c.Dora {
				if models.TileToID(&tile) == d {
					dora++
				}
			}
		}
		if c.Riichi || c.DoubleRiichi {
			for _, tile := range p.Tiles() {
				for _, d := range c.UraDora {
					if models.TileToID(&tile) == d {
						uradora++
					}
				}
			}
		}
		if akadora > 0 {
			yakuList = append(yakuList, yaku.AkaDora{Count: akadora})
		}
		if dora > 0 {
			yakuList = append(yakuList, yaku.Dora{Count: dora})
		}
		if uradora > 0 {
			yakuList = append(yakuList, yaku.UraDora{Count: uradora})
		}
		han += akadora + dora + uradora
		// if han is 5 or more, don't calculate fu
		slevel := HanToScoreLevel(han)
		if han > 4 {
			basicPoints := ScoreLevelToBasicPoints[slevel]
			if c.Jikaze == suits.Ton {
				score = 6 * basicPoints
			} else {
				score = 4 * basicPoints
			}
		} else {
			fu = CalculateFu(p, c)
			basicPoints := fu * int(math.Pow(2, float64(2+han)))
			if basicPoints >= 2000 {
				basicPoints = 2000
				slevel = Mangan
			}
			if c.Jikaze == suits.Ton {
				if c.Tsumo {
					score = 3 * roundUp(2*basicPoints, 100)
				} else {
					score = roundUp(6*basicPoints, 100)
				}
			} else {
				if c.Tsumo {
					score = 2*roundUp(basicPoints, 100) + roundUp(2*basicPoints, 100)
				} else {
					score = roundUp(4*basicPoints, 100)
				}
			}
		}
		return score, yakuList, nil, han, fu, ScoreLevelToString[slevel]
	}
}

/* WIP */

func CalculateHandShanten(h *models.Hand) int {
	return 0
}

func calculatePartitionShanten(p *models.Partition) int {
	return 0
}

func calculateStandardShanten(p *models.Partition) int {
	return 0
}

func calculateChiiToitsuShanten(p *models.Partition) int {
	return 0
}

func calculateKokushiShanten(p *models.Partition) int {
	return 0
}
