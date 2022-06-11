package main

import (
	"bufio"
	"fmt"
	"os"
	"riichi-calculator/src/models"
	"strings"
	"time"
)

func main() {
	var in string
	input := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter Hand> ")
		in, _ = input.ReadString('\n')
		hand, err := models.StringToHand(strings.TrimSpace(in))
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(hand)
			start := time.Now()
			found, partitions_with_waits := models.CheckTenpai(hand)
			end := time.Now()
			fmt.Printf("Found tenpai partitions in %f seconds.\n", end.Sub(start).Seconds())
			// fmt.Printf("Calculated %d partitions in %f seconds.\n", len(partitions), end.Sub(start).Seconds())
			if found {
				for wait, partitions := range partitions_with_waits {
					for _, p := range partitions {
						fmt.Printf("%s Waiting on %s\n", p.String(), models.TileToString[wait])
					}
				}
			} else {
				fmt.Println("The given hand is not in tenpai.")
			}
		}
	}
}
