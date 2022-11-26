package main

import "fmt"

func hurdleRace(k int32, height []int32) int32 {
	// Write your code here
	max := height[0]
	for _, n := range height {
		if n >= max {
			max = n
		}
	}

	if k > max {
		return 0
	}

	return (max - k)
}

func main() {
	fmt.Println(hurdleRace(4, []int32{1, 6, 3, 5, 2}))
	fmt.Println(hurdleRace(7, []int32{2, 5, 4, 5, 2}))
}
