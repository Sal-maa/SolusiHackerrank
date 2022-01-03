package main

import (
	"fmt"
	"math"
)

func diagonalDifference(arr [][]int32) int32 {
	// Write your code here
	var a, b int32

	for i := 0; i < len(arr); i++ {
		a += arr[i][i]
		b += arr[len(arr)-1-i][i]
	}
	return int32(math.Abs(float64(a) - float64(b)))
}

func main() {
	fmt.Println(diagonalDifference([][]int32{{11, 2, 4}, {4, 5, 6}, {10, 8, -12}}))
}
