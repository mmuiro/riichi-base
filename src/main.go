package main

import (
	"fmt"
	"riichi-calculator/src/models"
)

func main() {
	for {
		var in string
		fmt.Print("Enter Hand> ")
		fmt.Scanln(&in)
		hand, err := models.StringToHand(in)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(hand)
			partitions := models.CalculateAllPartitions(hand)
			fmt.Printf("Number of partitions: %d\n", len(partitions))
			for _, p := range partitions[:5] {
				fmt.Println(p)
			}
		}
	}
}
