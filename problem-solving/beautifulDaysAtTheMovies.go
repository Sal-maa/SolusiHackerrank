package main

import (
	"fmt"
	"math"
)

func beautifulDays(i int32, j int32, k int32) int32 {
	// Write your code here
	hasil := int32(0)
	days := make([]int32, j-i+1)
	revDays := make([]int32, j-i+1)
	days[0] = i

	for idx := range days {
		if idx != 0 {
			days[idx] = days[idx-1] + 1
		}
	}

	for idx, day := range days {
		res := int32(0)
		for day > 0 {
			remainder := day % 10
			res = (res * 10) + remainder
			day /= 10
		}
		revDays[idx] = res
	}

	for idx, res := range revDays {
		if int32(math.Abs(float64(days[idx]-res)))%k == 0 {
			hasil++
		}
	}

	return hasil
}

func main() {
	fmt.Println(beautifulDays(13, 45, 3))
	fmt.Println(beautifulDays(1, 2000000, 1000000000))
}
