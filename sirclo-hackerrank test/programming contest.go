package main

import (
	"fmt"
	"sort"
)

func minimizeBias(ratings []int32) int32 {
	// Write your code here

	sort.Slice(ratings, func(i, j int) bool {
		return ratings[i] < ratings[j]
	})
	fmt.Println(ratings)
	diff := []int32{}
	for i := 0; i < len(ratings)/2; i++ {
		x := int32(0)
		x = ratings[2*i+1] - ratings[2*i]
		diff = append(diff, int32(x))
	}

	sort.Slice(diff, func(i, j int) bool {
		return diff[i] < diff[j]
	})
	fmt.Println(diff)
	bias := int32(0)
	for _, v := range diff {
		bias += v
	}
	return bias
}

func main() {
	// fmt.Println(minimizeBias([]int32{2, 4, 5, 3, 7, 8}))
	fmt.Println(minimizeBias([]int32{10, 2, 5, 6, 10, 7, 6, 7, 9, 7, 7}))
}
