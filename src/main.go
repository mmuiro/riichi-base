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
			found, complete_partitions := models.CheckComplete(hand)
			end := time.Now()
			fmt.Printf("Found %d complete partitions in %f seconds.\n", len(complete_partitions), end.Sub(start).Seconds())
			// fmt.Printf("Calculated %d partitions in %f seconds.\n", len(partitions), end.Sub(start).Seconds())
			if found {
				for _, partition := range complete_partitions {
					fmt.Println(partition)
				}
			} else {
				fmt.Println("The given hand is incomplete.")
			}
		}
	}
}
