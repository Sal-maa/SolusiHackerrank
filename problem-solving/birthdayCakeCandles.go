package main

import (
	"fmt"
	"sort"
)

func birthdayCakeCandles(candles []int32) int32 {
	// Write your code here
	sort.Slice(candles, func(i, j int) bool {
		return candles[j] < candles[i]
	})

	max := candles[0]
	var a []int32
	for _, v := range candles {
		if v == max {
			a = append(a, v)
		}
	}
	return int32(len(a))
}

func main() {
	fmt.Println(birthdayCakeCandles([]int32{3, 2, 1, 3}))
}
