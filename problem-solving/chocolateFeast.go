package main

import "fmt"

func chocolateFeast(n int32, c int32, m int32) int32 {
	// Write your code here
	choco := n / c
	wrappers := n / c
	for wrappers >= m {
		choco += wrappers / m
		wrappers = wrappers/m + wrappers%m
	}
	return choco
}

func main() {
	fmt.Println(chocolateFeast(7, 3, 2))
	fmt.Println(chocolateFeast(75807, 357, 21088))
	fmt.Println(chocolateFeast(10, 2, 5))
	fmt.Println(chocolateFeast(20741, 60, 6676))
	fmt.Println(chocolateFeast(85582, 707, 73827))
	fmt.Println(chocolateFeast(38325, 35990, 17382))
	fmt.Println(chocolateFeast(90963, 40445, 40319))
	fmt.Println(chocolateFeast(20879, 13, 20651))
}
