package main

import "fmt"

func jumpingOnClouds(c []int32) int32 {
	// Write your code here
	var jump int32 = 0
	var i int = 0

	for i < len(c)-2 {
		if c[i+2] == 0 {
			i += 2
		} else {
			i += 1
		}
		jump++
	}

	if i == len(c)-2 {
		jump++
	}

	return jump
}

func main() {
	fmt.Println(jumpingOnClouds([]int32{0, 0, 1, 0, 0, 1, 0}))
	fmt.Println(jumpingOnClouds([]int32{0, 0, 1, 0, 0, 0, 0, 1, 0, 0}))
	fmt.Println(jumpingOnClouds([]int32{0, 1, 0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 1, 0, 0, 1, 0, 1, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 1, 0, 1, 0, 1, 0, 0, 0, 1, 0, 1, 0, 0, 0, 1, 0, 1, 0, 1, 0, 0, 0, 1, 0, 1, 0, 0, 0, 1, 0, 1, 0, 0, 0, 1, 0, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 0, 1, 0, 1, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0}))
	fmt.Println(jumpingOnClouds([]int32{0, 0}))
	fmt.Println(jumpingOnClouds([]int32{0, 1, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 1, 0, 1, 0, 0, 1, 0, 0, 0, 1, 0, 1, 0, 0, 0, 1, 0, 1, 0, 1, 0, 1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 1, 0, 1, 0, 0, 1, 0, 1, 0, 1, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 1, 0, 0, 1, 0, 1, 0}))
}
