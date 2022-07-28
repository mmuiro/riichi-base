package calculator

import (
	"fmt"
	"math"

	"github.com/mmuiro/riichi-base/src/models"
	"github.com/mmuiro/riichi-base/src/models/constants/groups"
	"github.com/mmuiro/riichi-base/src/models/constants/languages"
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

type InvalidAgariHandError struct {
	h *models.Hand
	w *models.Tile
}

func (e *InvalidAgariHandError) Error() string {
	return fmt.Sprintf("The hand %s is not complete with the tile %s.", e.h.String(), e.w.String())
}

type Score struct {
	Points            int
	WinningPartition  *models.Partition
	YakuList          []yaku.Yaku
	YakumanList       []yakuman.Yakuman
	YakumanMultiplier int
	Han               int
	Fu                int
	TsumoSplit        []int
	Scorelevel        ScoreLevel
}

func (s *Score) ScoreLevelName(l languages.Language) string {
	yakumanPrefixesByLanguage := [][]string{YakumanMultiplierNamesEN, YakumanMultiplierNamesJA}
	scoreLevelsByLanguage := []map[ScoreLevel]string{ScoreLevelToStringEN, ScoreLevelToStringJA}
	if s.YakumanMultiplier > 0 {
		var base string
		if l == languages.EN {
			base = "Yakuman"
		} else {
			base = "役満"
		}
		return yakumanPrefixesByLanguage[l][s.YakumanMultiplier-1] + base
	} else if s.Points > 0 {
		return scoreLevelsByLanguage[l][s.Scorelevel]
	}
	return ""
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
	score, err := CalculateScoreVerbose(h, w, c)
	if err != nil {
		return 0, err
	}
	return score.Points, nil
}

func CalculateScoreVerbose(h *models.Hand, w *models.Tile, c *yaku.Conditions) (*Score, error) {
	if w == nil {
		return &Score{Points: 0}, &NoAgariError{h: h}
	}
	if !h.Tenpai {
		err := h.RemoveTile(w)
		if err != nil {
			return &Score{Points: 0}, &InvalidAgariHandError{h: h, w: w}
		}
		var waitMaps [][]int
		var tenpaiPartitions []models.Partition
		h.Tenpai, tenpaiPartitions, waitMaps = models.CheckTenpai(h)
		if !h.Tenpai {
			return &Score{Points: 0}, &NoAgariError{h: h}
		}
		models.AssignWaitMap(h, tenpaiPartitions, waitMaps)
	}
	agari, partitions := models.CheckAgari(h, w, c.Tsumo)
	if !agari {
		return &Score{Points: 0}, &NoAgariError{h: h}
	}
	maxPoints := 0
	var bestScore *Score
	for _, p := range partitions {
		score := calculatePartitionScore(&p, c)
		if score.Points > maxPoints {
			bestScore = score
			maxPoints = score.Points
		}
	}
	if maxPoints == 0 {
		return &Score{Points: 0}, &NoYakuError{h: h}
	}
	return bestScore, nil
}

func calculatePartitionScore(p *models.Partition, c *yaku.Conditions) *Score {
	yakumanMultiplier, yakumanList := FindYakuman(p, c)
	points := 0
	var tsumoSplit []int
	// has a yakuman
	if yakumanMultiplier > 0 {
		if c.Jikaze == suits.Ton {
			points = 6 * ScoreLevelToBasicPoints[Yakuman]
		} else {
			points = 4 * ScoreLevelToBasicPoints[Yakuman]
		}
		if c.Tsumo {
			if c.Jikaze == suits.Ton {
				tsumoSplit = []int{yakumanMultiplier * points / 3}
			} else {
				tsumoSplit = []int{yakumanMultiplier * points / 4, yakumanMultiplier * points / 2}
			}
		}
		return &Score{Points: yakumanMultiplier * points, WinningPartition: p, YakumanList: yakumanList, YakumanMultiplier: yakumanMultiplier, Scorelevel: Yakuman, TsumoSplit: tsumoSplit}
	} else {
		// Find the han
		han, yakuList := FindHanAndYaku(p, c)
		var fu int
		if han == 0 {
			return &Score{Points: 0}
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
		slevel := HanToScoreLevel(han)
		fu = CalculateFu(p, c)
		if han > 4 {
			basicPoints := ScoreLevelToBasicPoints[slevel]
			if c.Jikaze == suits.Ton {
				points = 6 * basicPoints
			} else {
				points = 4 * basicPoints
			}
		} else {
			basicPoints := fu * int(math.Pow(2, float64(2+han)))
			if basicPoints >= 2000 {
				basicPoints = 2000
				slevel = Mangan
			}
			if c.Jikaze == suits.Ton {
				if c.Tsumo {
					points = 3 * roundUp(2*basicPoints, 100)
					tsumoSplit = []int{points / 3}
				} else {
					points = roundUp(6*basicPoints, 100)
				}
			} else {
				if c.Tsumo {
					points = 2*roundUp(basicPoints, 100) + roundUp(2*basicPoints, 100)
					tsumoSplit = []int{roundUp(basicPoints, 100), roundUp(2*basicPoints, 100)}
				} else {
					points = roundUp(4*basicPoints, 100)
				}
			}
		}
		return &Score{Points: points, WinningPartition: p, YakuList: yakuList, Han: han, Fu: fu, Scorelevel: slevel, TsumoSplit: tsumoSplit}
	}
}

// Returns a slice of tile IDs corresponding to the tiles the hand is waiting on, if it is in tenpai.
func CalculateWaitTiles(h *models.Hand) []int {
	tenpai, _, waitLists := models.CheckTenpai(h)
	waits := make([]int, 0)
	hits := make([]bool, 34)
	if tenpai {
		for _, waitList := range waitLists {
			for _, tileID := range waitList {
				if !hits[tileID] {
					waits = append(waits, tileID)
					hits[tileID] = true
				}
			}
		}
	}
	return waits
}

// Returns a map from tileIDs of discardable tiles to slices of tileIDs of the waits they produce.
func CalculateWaitTilesFull(h *models.Hand) map[int][]int {
	discToWait := make(map[int][]int)
	for _, tile := range h.ClosedTiles {
		h.RemoveTile(&tile)
		discToWait[models.TileToID(&tile)] = CalculateWaitTiles(h)
		h.AddTile(&tile)
	}
	return discToWait
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
