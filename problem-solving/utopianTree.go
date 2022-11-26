package main

import "fmt"

func utopianTree(n int32) int32 {
	// Write your code here
	init := int32(1)
	trigger := make([]int, n)

	for idx := range trigger {
		if idx%2 == 0 {
			init = init * 2
		} else {
			init++
		}
	}

	return init
}

func main() {
	fmt.Println(utopianTree(4))
	fmt.Println(utopianTree(7))
}
