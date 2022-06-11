package models

import (
	"riichi-calculator/src/models/constants/groups"
	"riichi-calculator/src/models/constants/waits"
	"riichi-calculator/src/utils"
	"sort"
	"strings"
)

type Partition struct {
	Mentsu       []Mentsu
	mentsuCounts map[groups.MentsuType]int
	Wait         waits.WaitKind
}

func (p Partition) String() string {
	mentsuStrings := make([]string, len(p.Mentsu))
	for i, mentsu := range p.Mentsu {
		mentsuStrings[i] = mentsu.String()
	}
	return strings.Join(mentsuStrings, " -")
}

func (p Partition) Tiles() []Tile {
	tiles := make([]Tile, 0)
	for _, mentsu := range p.Mentsu {
		tiles = append(tiles, mentsu.tiles...)
	}
	return tiles
}

func (p Partition) TileCount() int {
	sum := 0
	for _, p := range p.Mentsu {
		sum += len(p.tiles)
	}
	return sum
}

func CalculateAllPartitions(h *Hand) []Partition {
	results := make([]Partition, 0)
	nonMeldTiles := make([]Tile, 0)
	// buggy. needs to exclude it from the meld after it is found once.
	for _, tile := range h.tiles {
		nonMeld := true
		for _, meld := range h.melds {
			for _, meldTile := range meld.tiles {
				if meldTile.Equals(&tile) {
					nonMeld = false
				}
			}
		}
		if nonMeld {
			nonMeldTiles = append(nonMeldTiles, tile)
		}
	}
	memo := make(map[string][][]Mentsu)
	for _, partition := range calculatePartitionsFromTiles(nonMeldTiles, memo) {
		newPartition := Partition{Mentsu: append(partition, h.melds...)}
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

func checkAndAssignMentsuCounts(p *Partition) {
	if p.mentsuCounts == nil {
		p.mentsuCounts = make(map[groups.MentsuType]int)
		for _, mentsu := range p.Mentsu {
			p.mentsuCounts[mentsu.kind]++
		}
	}
}

// Checks whether the given hand partition has 1 pair and 4 other complete groups (sets and sequences).
func CheckStandard(p *Partition) bool {
	checkAndAssignMentsuCounts(p)
	return len(p.Mentsu) == 5 && p.mentsuCounts[groups.Toitsu] == 1 &&
		p.mentsuCounts[groups.Shuntsu]+p.mentsuCounts[groups.Koutsu]+p.mentsuCounts[groups.Kantsu] == 4
}

// Checks whether the given hand partition has Chii Toitsu (7 pairs).
func CheckChiiToitsu(p *Partition) bool {
	return len(p.Mentsu) == 7 && utils.All(utils.FuncMap(func(m Mentsu) bool {
		return m.kind == groups.Toitsu
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

// Tenpai Checks

func CheckRyanmen(p *Partition) bool {
	checkAndAssignMentsuCounts(p)
	return len(p.Mentsu) == 5 && p.mentsuCounts[groups.Toitsu] == 1 &&
		p.mentsuCounts[groups.Shuntsu]+p.mentsuCounts[groups.Koutsu]+p.mentsuCounts[groups.Kantsu] == 3 &&
		p.mentsuCounts[groups.Ryanmen] == 1
}

func CheckKanchan(p *Partition) bool {
	checkAndAssignMentsuCounts(p)
	return len(p.Mentsu) == 5 && p.mentsuCounts[groups.Toitsu] == 1 &&
		p.mentsuCounts[groups.Shuntsu]+p.mentsuCounts[groups.Koutsu]+p.mentsuCounts[groups.Kantsu] == 3 &&
		p.mentsuCounts[groups.Kanchan] == 1
}

func CheckPenchan(p *Partition) bool {
	checkAndAssignMentsuCounts(p)
	return len(p.Mentsu) == 5 && p.mentsuCounts[groups.Toitsu] == 1 &&
		p.mentsuCounts[groups.Shuntsu]+p.mentsuCounts[groups.Koutsu]+p.mentsuCounts[groups.Kantsu] == 3 &&
		p.mentsuCounts[groups.Penchan] == 1
}

func CheckShanpon(p *Partition) bool {
	checkAndAssignMentsuCounts(p)
	return len(p.Mentsu) == 5 &&
		p.mentsuCounts[groups.Shuntsu]+p.mentsuCounts[groups.Koutsu]+p.mentsuCounts[groups.Kantsu] == 3 &&
		p.mentsuCounts[groups.Toitsu] == 2
}

func CheckTanki(p *Partition) bool {
	checkAndAssignMentsuCounts(p)
	return len(p.Mentsu) == 5 &&
		p.mentsuCounts[groups.Shuntsu]+p.mentsuCounts[groups.Koutsu]+p.mentsuCounts[groups.Kantsu] == 4
}

func CheckKokushiSingle(p *Partition) bool {
	// it is a bit inefficient.
	if len(p.Tiles()) != 13 || len(p.Mentsu) != 12 {
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
	missing := 0
	for _, id := range KokushiTileIDs {
		if !includedKokushiTiles[id] {
			missing += 1
		}
	}
	return missing == 1
}
