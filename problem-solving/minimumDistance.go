package main

import "fmt"

func minimumDistances(a []int32) int32 {
	// Write your code here
	mapInit := make(map[int32]int)
	min := int32(len(a) - 1)
	change := false
	for k, v := range a {
		if _, ok := mapInit[v]; !ok {
			mapInit[v] = k
		} else {
			val := int32(k - mapInit[v])
			if val <= min {
				min = val
				change = true
			}
		}
	}

	if !change {
		return -1
	}

	return min
}

func main() {
	fmt.Println(minimumDistances([]int32{7, 1, 3, 4, 1, 7}))
	fmt.Println(minimumDistances([]int32{1, 2, 3, 4, 10}))
	fmt.Println(minimumDistances([]int32{10}))
}
