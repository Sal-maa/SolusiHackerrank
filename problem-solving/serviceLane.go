package main

import "fmt"

func serviceLane(width []int32, cases [][]int32) []int32 {
	// Write your code here
	pass := []int32{}
	for i := 0; i < len(cases); i++ {
		vehicles := width[cases[i][0] : cases[i][1]+1]
		min := vehicles[0]
		for _, v := range vehicles {
			if v <= min {
				min = v
			}
		}
		pass = append(pass, min)
	}
	return pass
}

func main() {
	fmt.Println(serviceLane([]int32{2, 3, 1, 2, 3, 2, 3, 3}, [][]int32{{0, 3}, {4, 6}, {6, 7}, {3, 5}, {0, 7}}))
	fmt.Println(serviceLane([]int32{1, 2, 2, 2, 1}, [][]int32{{2, 3}, {1, 4}, {2, 4}, {2, 4}, {2, 4}}))
}
