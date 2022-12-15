package main

import (
	"fmt"
	"sort"
)

func equalizeArray(arr []int32) int32 {
	// Write your code here
	mapInput := make(map[int32]int)
	var count int32
	for _, a := range arr {
		mapInput[a]++
	}

	var countArr []int
	for _, v := range mapInput {
		countArr = append(countArr, v)
	}

	if len(countArr) == 1 {
		return 0
	}

	sort.Ints(countArr)
	for i, v := range countArr {
		if i < len(countArr)-1 {
			count += int32(v)
		}
	}

	return count
}

func main() {
	fmt.Println(equalizeArray([]int32{10, 27, 9, 10, 100, 38, 30, 32, 45, 29, 27, 29, 32, 38, 32, 38, 14, 38, 29, 30, 63, 29, 63, 91, 54, 10, 63}))
	fmt.Println(equalizeArray([]int32{37, 32, 97, 35, 76, 62}))
	fmt.Println(equalizeArray([]int32{51, 51, 51, 51}))
	fmt.Println(equalizeArray([]int32{24, 29, 70, 43, 12, 27, 29, 24, 41, 12, 41, 43, 24, 70, 24, 100, 41, 43, 43, 100, 29, 70, 100, 43, 41, 27, 70, 70, 59, 41, 24, 24, 29, 43, 24, 27, 70, 24, 27, 70, 24, 70, 27, 24, 43, 27, 100, 41, 12, 70, 43, 70, 62, 12, 59, 29, 62, 41, 100, 43, 43, 59, 59, 70, 12, 27, 43, 43, 27, 27, 27, 24, 43, 43, 62, 43, 70, 29}))
}
