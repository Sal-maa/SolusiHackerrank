package main

import "fmt"

func permutationEquation(p []int32) []int32 {
	// Write your code here
	mapFunc := make(map[int32]int32, 0)

	for idx, val := range p {
		mapFunc[val] = int32(idx) + 1
	}

	var res []int32
	for i := 1; i <= len(p); i++ {
		res = append(res, mapFunc[mapFunc[int32(i)]])
	}

	return res
}

func main() {
	fmt.Println(permutationEquation([]int32{1, 4, 2, 3, 5}))
	fmt.Println(permutationEquation([]int32{2, 5, 11, 10, 1, 14, 7, 3, 16, 9, 8, 6, 18, 12, 15, 17, 13, 4}))
}
