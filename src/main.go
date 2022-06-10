package main

import (
	"bufio"
	"fmt"
	"math"
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
			partitions := models.CalculateAllPartitions(hand)
			end := time.Now()
			fmt.Printf("Calculated %d partitions in %f seconds.\n", len(partitions), end.Sub(start).Seconds())
			for i := 0; i < int(math.Min(float64(len(partitions)), 5)); i++ {
				fmt.Println(partitions[i])
			}
		}
	}
}
