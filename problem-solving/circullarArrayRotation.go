package main

import "fmt"

func circularArrayRotation(a []int32, k int32, queries []int32) []int32 {
	// Write your code here
	var rotations int32 = k % int32(len(a))
	var result []int32 = []int32{}

	for _, index := range queries {
		index = (index - rotations + int32(len(a))) % int32(len(a))
		result = append(result, a[index])
	}

	return result
}

func main() {
	fmt.Println(circularArrayRotation([]int32{1, 2, 3}, 2, []int32{0, 1, 2}))
	fmt.Println(circularArrayRotation([]int32{3, 4, 5}, 2, []int32{1, 2}))
}
