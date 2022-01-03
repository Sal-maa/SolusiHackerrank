package main

import (
	"fmt"
	"math"
	"sort"
)

func minDiff(arr []int32) int32 {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	angka := int32(0)
	for i := 0; i < len(arr)-1; i++ {
		angka = angka + int32(math.Abs(float64(arr[i]-arr[i+1])))
		fmt.Println(arr[i], arr[i+1])
	}
	return angka
}

func main() {
	fmt.Println(minDiff([]int32{5, 1, 3, 7, 3}))
}
