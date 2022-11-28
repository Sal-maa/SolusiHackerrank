package main

import "fmt"

func jumpingOnClouds(c []int32, k int32) int32 {
	e := int32(100)
	var count int
	for i := int(k); count < 1; i += int(k) {
		if i%len(c) == 0 {
			count++
		}
		if c[i%len(c)] == 1 {
			e -= 2
		}
		e--
	}
	return e
}

func main() {
	fmt.Println(jumpingOnClouds([]int32{0, 0, 1, 0, 0, 1, 1, 0}, 2))
	fmt.Println(jumpingOnClouds([]int32{1, 1, 1, 0, 1, 1, 0, 0, 0, 0}, 3))
}
