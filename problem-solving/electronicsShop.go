package main

import "fmt"

func getMoneySpent(keyboards []int32, drives []int32, b int32) int32 {
	/*
	 * Write your code here.
	 */
	var totalMax int32

	for _, k := range keyboards {
		for _, d := range drives {
			if k+d > totalMax && k+d <= b {
				totalMax = k + d
			}
		}

	}

	if totalMax == 0 {
		totalMax = -1
	}
	return totalMax
}

func main() {
	fmt.Println(getMoneySpent([]int32{3, 1}, []int32{5, 2, 8}, 10))
	fmt.Println(getMoneySpent([]int32{4}, []int32{5}, 5))
}
