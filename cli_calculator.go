package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
	"time"

	"github.com/mmuiro/riichi-base/src/calculator"
	"github.com/mmuiro/riichi-base/src/models"
	"github.com/mmuiro/riichi-base/src/models/constants/suits"
	"github.com/mmuiro/riichi-base/src/models/yaku"
	"github.com/mmuiro/riichi-base/src/utils"
)

func main() {
	var in string
	input := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter Hand> ")
		in, _ = input.ReadString('\n')
		hand, lastTile, err := models.StringToHand(strings.TrimSpace(in))
		if err != nil {
			fmt.Println(err)
		} else {
			hand.RemoveTile(lastTile)
			fmt.Println(hand)
			start := time.Now()
			if found, tenpais, waitLists := models.CheckTenpai(hand); found {
				c := yaku.Conditions{Menzenchin: false, Jikaze: suits.Ton, Bakaze: suits.Ton, Tsumo: true, Dora: []int{models.SuitAndValueToID(suits.Man, 1)}, UraDora: []int{models.SuitAndValueToID(suits.Chun, 0)}} // should be customizable
				models.AssignWaitMap(hand, tenpais, waitLists)
				hand.Tenpai = true
				if lastTile != nil {
					// calculate the score, if possible (hand has agari)
					var score *calculator.Score
					score, err = calculator.CalculateScoreVerbose(hand, lastTile, &c)
					if err != nil {
						fmt.Println(err)
						continue
					} else {
						fmt.Println(score.WinningPartition)
						if len(score.YakumanList) > 0 {
							for _, y := range score.YakumanList {
								fmt.Printf("%s\n", y.Name())
							}
						} else {
							for _, y := range score.YakuList {
								fmt.Printf("%s - %d han\n", y.Name(), y.Han(!c.Menzenchin))
							}
							fmt.Printf("%d han", score.Han)
							if score.Fu > 0 {
								fmt.Printf(" %d fu", score.Fu)
							}
							fmt.Println()
						}
						fmt.Println("--------------------------")
						pre := ""
						if c.Jikaze == suits.Ton {
							pre += "Dealer "
						}
						if c.Tsumo {
							pre += "Tsumo"
						} else {
							pre += "Ron"
						}
						fmt.Println(pre)
						if score.ScoreLevel != "" {
							fmt.Printf("%s - %d pts\n", score.ScoreLevel, score.Points)
						} else {
							fmt.Printf("%d pts\n", score.Points)
						}
					}
				} else {
					fmt.Println("The given hand is in tenpai, with the following waits:")
					for i, p := range tenpais {
						waitTiles := utils.FuncMap(func(id int) string {
							return models.TileToString[id]
						}, waitLists[i])
						fmt.Printf("%s, waiting on %s\n", p.String(), strings.Join(waitTiles, " "))
					}
				}
			} else {
				fmt.Println("The given hand is not in tenpai, nor is it complete.")
				partitions := models.CalculateAllPartitions(hand)
				for i := 0; i < int(math.Min(5, float64(len(partitions)))); i++ {
					fmt.Println(partitions[i])
				}
			}
			end := time.Now()
			fmt.Printf("Calculation done in %f seconds.\n", end.Sub(start).Seconds())
			// fmt.Printf("Calculated %d partitions in %f seconds.\n", len(partitions), end.Sub(start).Seconds())
		}
	}
}
