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
	"github.com/mmuiro/riichi-base/src/models/yaku/yakuman"
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
			fmt.Println(hand)
			start := time.Now()
			if lastTile != nil {
				err = hand.RemoveTile(lastTile)
				if err != nil {
					fmt.Println(err)
					continue
				}
			}
			if found, tenpais, waitLists := models.CheckTenpai(hand); found {
				models.AssignWaitMap(hand, tenpais, waitLists)
				hand.Tenpai = true
				if lastTile != nil {
					// calculate the score, if possible (hand has agari)
					c := yaku.Conditions{Menzenchin: false, Jikaze: suits.Nan, Bakaze: suits.Nan, Tsumo: false, Dora: []int{models.SuitAndValueToID(suits.Chun, 0)}} // should be customizable
					var score int
					var best *models.Partition
					var yakuList []yaku.Yaku
					var yakumanList []yakuman.Yakuman
					var slevel string
					var han, fu int
					score, best, yakuList, yakumanList, han, fu, slevel, err = calculator.CalculateScoreVerbose(hand, lastTile, &c)
					if err != nil {
						fmt.Println(err)
						continue
					} else {
						fmt.Println(best)
						if len(yakumanList) > 0 {
							for _, y := range yakumanList {
								fmt.Printf("%s\n", y.Name())
							}
						} else {
							for _, y := range yakuList {
								fmt.Printf("%s - %d han\n", y.Name(), y.Han(!c.Menzenchin))
							}
							fmt.Printf("%d han", han)
							if fu > 0 {
								fmt.Printf(" %d fu", fu)
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
						if slevel != "" {
							fmt.Printf("%s - %d pts\n", slevel, score)
						} else {
							fmt.Printf("%d pts\n", score)
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
