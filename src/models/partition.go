package models

import (
	"sort"
	"strings"

	"github.com/mmuiro/riichi-base/src/models/constants/groups"
	"github.com/mmuiro/riichi-base/src/models/constants/waits"
	"github.com/mmuiro/riichi-base/src/utils"
)

type Partition struct {
	Mentsu       []Mentsu
	MentsuCounts map[groups.MentsuType]int
	Wait         waits.WaitKind
}

func (p Partition) String() string {
	mentsuStrings := make([]string, len(p.Mentsu))
	for i, mentsu := range p.Mentsu {
		mentsuStrings[i] = mentsu.String()
	}
	return strings.Join(mentsuStrings, " -")
}

func (p *Partition) Tiles() []Tile {
	tiles := make([]Tile, 0)
	for _, mentsu := range p.Mentsu {
		tiles = append(tiles, mentsu.Tiles...)
	}
	return tiles
}

func (p *Partition) TileCount() int {
	sum := 0
	for _, p := range p.Mentsu {
		sum += len(p.Tiles)
	}
	return sum
}

func CalculateAllPartitions(h *Hand) []Partition {
	results := make([]Partition, 0)
	nonMeldTiles := make([]Tile, len(h.Tiles))
	copy(nonMeldTiles, h.Tiles)
	for _, meld := range h.Melds {
		for _, tile := range meld.Tiles {
			if i := getTileIndex(nonMeldTiles, tile.Suit, tile.Value); i > -1 {
				nonMeldTiles = utils.RemoveIndex(nonMeldTiles, i)
			}
		}
	}
	memo := make(map[string][][]Mentsu)
	for _, partition := range calculatePartitionsFromTiles(nonMeldTiles, memo) {
		newPartition := Partition{Mentsu: append(partition, h.Melds...)}
		results = append(results, newPartition)
	}
	sort.Slice(results, func(i, j int) bool {
		return len(results[i].Mentsu) < len(results[j].Mentsu)
	})
	return results
}

func calculatePartitionsFromTiles(rest []Tile, memo map[string][][]Mentsu) [][]Mentsu {
	results := make([][]Mentsu, 0)
	if len(rest) == 0 {
		results = append(results, []Mentsu{})
		return results
	}
	key := TilesToString(rest)
	if memoResult, ok := memo[key]; ok {
		return memoResult
	}
	currentTile := rest[0]
	// use first tile as a single
	nextRest := utils.RemoveIndex(utils.Clone(rest), 0)
	singleMentsu, _ := CreateMentsu([]Tile{currentTile}, false)
	for _, partition := range calculatePartitionsFromTiles(nextRest, memo) {
		results = append(results, append(partition, *singleMentsu))
	}
	// try to create a pair
	if index := getTileIndex(nextRest, currentTile.Suit, currentTile.Value); index >= 0 {
		var pairTile Tile
		results, nextRest, pairTile = removeAndGetPartitions(results, nextRest, index, []Tile{currentTile}, memo)
		// try to create a triplet (no need for quads, as those are melds)
		if index = getTileIndex(nextRest, currentTile.Suit, currentTile.Value); index >= 0 {
			results, _, _ = removeAndGetPartitions(results, nextRest, index, []Tile{currentTile, pairTile}, memo)
		}
	}
	// try to create a side-wait/sequence
	nextRest = utils.RemoveIndex(utils.Clone(rest), 0)
	if currentTile.Value > 0 && currentTile.Value < 9 && !currentTile.IsHonor() {
		if index := getTileIndex(nextRest, currentTile.Suit, currentTile.Value+1); index >= 0 {
			var secondTile Tile
			results, nextRest, secondTile = removeAndGetPartitions(results, nextRest, index, []Tile{currentTile}, memo)
			if index := getTileIndex(nextRest, currentTile.Suit, currentTile.Value+2); currentTile.Value < 8 && index >= 0 {
				results, _, _ = removeAndGetPartitions(results, nextRest, index, []Tile{currentTile, secondTile}, memo)
			}
		}
	}
	// try to create a closed-wait
	nextRest = utils.RemoveIndex(utils.Clone(rest), 0)
	if currentTile.Value > 0 && currentTile.Value < 8 && !currentTile.IsHonor() {
		if index := getTileIndex(nextRest, currentTile.Suit, currentTile.Value+2); index >= 0 {
			results, _, _ = removeAndGetPartitions(results, nextRest, index, []Tile{currentTile}, memo)
		}
	}
	copy(memo[key], results)
	return results
}

func removeAndGetPartitions(results [][]Mentsu, rest []Tile, index int, mentsuTiles []Tile, memo map[string][][]Mentsu) ([][]Mentsu, []Tile, Tile) {
	nextTile := rest[index]
	rest = utils.RemoveIndex(rest, index)
	mentsu, _ := CreateMentsu(append(mentsuTiles, nextTile), false)
	for _, partition := range calculatePartitionsFromTiles(rest, memo) {
		results = append(results, append(partition, *mentsu))
	}
	return results, rest, nextTile
}

// Hand Completion Checks

func CheckAndAssignMentsuCounts(p *Partition) {
	if p.MentsuCounts == nil {
		AssignMentsuCounts(p)
	}
}

func AssignMentsuCounts(p *Partition) {
	p.MentsuCounts = make(map[groups.MentsuType]int)
	for _, mentsu := range p.Mentsu {
		p.MentsuCounts[mentsu.Kind]++
	}
}

// Checks whether the given hand partition has 1 pair and 4 other complete groups (sets and sequences).
func CheckStandard(p *Partition) bool {
	CheckAndAssignMentsuCounts(p)
	return len(p.Mentsu) == 5 && p.MentsuCounts[groups.Toitsu] == 1 &&
		p.MentsuCounts[groups.Shuntsu]+p.MentsuCounts[groups.Koutsu]+p.MentsuCounts[groups.Kantsu] == 4
}

// Checks whether the given hand partition has Chii Toitsu (7 pairs).
func CheckChiiToitsu(p *Partition) bool {
	return len(p.Mentsu) == 7 && utils.All(utils.FuncMap(func(m Mentsu) bool {
		return m.Kind == groups.Toitsu
	}, p.Mentsu))
}

// Checks whether the given hand partition has Kokushi Musou (13 orphans).
func CheckKokushi(p *Partition) bool {
	// it is a bit inefficient.
	if len(p.Tiles()) != 14 || len(p.Mentsu) != 13 {
		return false
	}
	includedKokushiTiles := make(map[int]bool)
	for _, tile := range p.Tiles() {
		id := TileToID(&tile)
		if !utils.Contains(KokushiTileIDs, id) {
			return false
		}
		includedKokushiTiles[id] = true
	}
	for _, id := range KokushiTileIDs {
		if !includedKokushiTiles[id] {
			return false
		}
	}
	return true
}

/*
Tenpai Checks
These functions will check if the partition is in a kind of wait;
if it is, then it also returns the waiting tiles' ids.
*/

func CheckRyanmen(p *Partition) (bool, []int) {
	CheckAndAssignMentsuCounts(p)
	cond := len(p.Mentsu) == 5 && p.MentsuCounts[groups.Toitsu] == 1 &&
		p.MentsuCounts[groups.Shuntsu]+p.MentsuCounts[groups.Koutsu]+p.MentsuCounts[groups.Kantsu] == 3 &&
		p.MentsuCounts[groups.Ryanmen] == 1
	if cond {
		for _, mentsu := range p.Mentsu {
			if mentsu.Kind == groups.Ryanmen {
				waits := []int{TileToID(&mentsu.Tiles[0]) - 1, TileToID(&mentsu.Tiles[1]) + 1}
				return cond, waits
			}
		}
	}
	return cond, nil
}

func CheckKanchan(p *Partition) (bool, []int) {
	CheckAndAssignMentsuCounts(p)
	cond := len(p.Mentsu) == 5 && p.MentsuCounts[groups.Toitsu] == 1 &&
		p.MentsuCounts[groups.Shuntsu]+p.MentsuCounts[groups.Koutsu]+p.MentsuCounts[groups.Kantsu] == 3 &&
		p.MentsuCounts[groups.Kanchan] == 1
	if cond {
		for _, mentsu := range p.Mentsu {
			if mentsu.Kind == groups.Kanchan {
				waits := []int{TileToID(&mentsu.Tiles[0]) + 1}
				return cond, waits
			}
		}
	}
	return cond, nil
}

func CheckPenchan(p *Partition) (bool, []int) {
	CheckAndAssignMentsuCounts(p)
	cond := len(p.Mentsu) == 5 && p.MentsuCounts[groups.Toitsu] == 1 &&
		p.MentsuCounts[groups.Shuntsu]+p.MentsuCounts[groups.Koutsu]+p.MentsuCounts[groups.Kantsu] == 3 &&
		p.MentsuCounts[groups.Penchan] == 1
	if cond {
		for _, mentsu := range p.Mentsu {
			if mentsu.Kind == groups.Penchan {
				var wait int
				if mentsu.Tiles[0].Value == 1 {
					wait = TileToID(&mentsu.Tiles[0]) + 2
				} else {
					wait = TileToID(&mentsu.Tiles[0]) - 1
				}
				return cond, []int{wait}
			}
		}
	}
	return cond, nil
}

func CheckShanpon(p *Partition) (bool, []int) {
	CheckAndAssignMentsuCounts(p)
	cond := len(p.Mentsu) == 5 &&
		p.MentsuCounts[groups.Shuntsu]+p.MentsuCounts[groups.Koutsu]+p.MentsuCounts[groups.Kantsu] == 3 &&
		p.MentsuCounts[groups.Toitsu] == 2
	if cond {
		waits := make([]int, 0)
		for _, mentsu := range p.Mentsu {
			if mentsu.Kind == groups.Toitsu {
				waits = append(waits, TileToID(&mentsu.Tiles[0]))
			}
		}
		return cond, waits
	}
	return cond, nil
}

func CheckTanki(p *Partition) (bool, []int) {
	CheckAndAssignMentsuCounts(p)
	cond := (len(p.Mentsu) == 5 &&
		p.MentsuCounts[groups.Shuntsu]+p.MentsuCounts[groups.Koutsu]+p.MentsuCounts[groups.Kantsu] == 4 &&
		p.MentsuCounts[groups.Single] == 1) ||
		(len(p.Mentsu) == 7 && p.MentsuCounts[groups.Toitsu] == 6 && p.MentsuCounts[groups.Single] == 1)
	if cond {
		for _, mentsu := range p.Mentsu {
			if mentsu.Kind == groups.Single {
				return cond, []int{TileToID(&mentsu.Tiles[0])}
			}
		}
	}
	return cond, nil
}

func CheckKokushiSingle(p *Partition) (bool, []int) {
	// it is a bit inefficient.
	if len(p.Tiles()) != 13 || len(p.Mentsu) != 12 {
		return false, nil
	}
	includedKokushiTiles := make(map[int]bool)
	for _, tile := range p.Tiles() {
		id := TileToID(&tile)
		if !utils.Contains(KokushiTileIDs, id) {
			return false, nil
		}
		includedKokushiTiles[id] = true
	}
	missing, wait := 0, 0
	for _, id := range KokushiTileIDs {
		if !includedKokushiTiles[id] {
			missing += 1
			wait = id
		}
	}
	if missing == 1 {
		return true, []int{wait}
	}
	return false, nil
}

func CheckKokushiThirteen(p *Partition) (bool, []int) {
	if len(p.Tiles()) != 13 || len(p.Mentsu) != 13 {
		return false, nil
	}
	includedKokushiTiles := make(map[int]bool)
	for _, tile := range p.Tiles() {
		id := TileToID(&tile)
		if !utils.Contains(KokushiTileIDs, id) {
			return false, nil
		}
		includedKokushiTiles[id] = true
	}

	for _, id := range KokushiTileIDs {
		if !includedKokushiTiles[id] {
			return false, nil
		}
	}
	waits := make([]int, 13)
	copy(waits, KokushiTileIDs)
	return true, waits
}

func CheckJunseiChuuren(p *Partition) bool {
	if len(p.Tiles()) != 13 {
		return false
	}
	suit := p.Mentsu[0].Suit
	for _, mentsu := range p.Mentsu {
		if mentsu.Open || mentsu.Suit != suit {
			return false
		}
	}
	tileCounts := make([]int, 9)
	for _, tile := range p.Tiles() {
		tileCounts[tile.Value-1]++
	}
	if tileCounts[0] != 3 || tileCounts[8] != 3 {
		return false
	}
	for i := 1; i < 8; i++ {
		if tileCounts[i] != 1 {
			return false
		}
	}
	return true
}
