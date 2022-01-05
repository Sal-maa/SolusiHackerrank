package main

import (
	"fmt"
	"sort"
)

func miniMaxSum(arr []int32) {
	// Write your code here
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	var sum int64
	for i := 0; i < 5; i++ {
		sum += int64(arr[i])
	}

	fmt.Print(sum-int64(arr[4]), sum-int64(arr[0]))

}

func main() {
	miniMaxSum([]int32{1, 2, 3, 4, 5})
}
