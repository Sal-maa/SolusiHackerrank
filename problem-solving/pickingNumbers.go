package main

import (
	"fmt"
	"sort"
)

func pickingNumbers(a []int32) int32 {
	// Write your code here
	var count, result int32
	mapNumber := make(map[int32]int32)

	for _, v := range a {
		mapNumber[v]++
	}

	unique := make([]int, 0)
	for k, v := range mapNumber {
		unique = append(unique, int(k))
		if v > result {
			result = v
		}
	}

	sort.Ints(unique)

	for i := 0; i < len(unique)-1; i++ {
		if (unique[1+i] - unique[i]) <= 1 {
			count = mapNumber[int32(unique[1+i])] + mapNumber[int32(unique[i])]
			if count > result {
				result = count
			}
		}
	}

	return result
}

func main() {
	fmt.Println(pickingNumbers([]int32{4, 6, 5, 3, 3, 1}))
	fmt.Println(pickingNumbers([]int32{1, 2, 2, 3, 1, 2}))
}
