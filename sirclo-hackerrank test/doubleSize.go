package main

import (
	"fmt"
	"sort"
)

func doubleSize(arr []int64, b int64) int64 {
	// Write your code here
	sort.SliceStable(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	for i := int64(0); i < int64(len(arr)); i++ {
		if arr[i] == b {
			b = 2 * b
		}
	}
	return b
}

func main() {
	fmt.Println(doubleSize([]int64{1, 2, 3, 1, 2}, 1))
}
