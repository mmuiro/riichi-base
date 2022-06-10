package models

import (
	"riichi-calculator/src/utils"
	"sort"
	"strings"
)

type Partition struct {
	Mentsu []Mentsu
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
		newPartition := Partition{append(partition, h.melds...)}
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
