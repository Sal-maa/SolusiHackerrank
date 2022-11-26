package main

import "fmt"

func angryProfessor(k int32, a []int32) string {
	// Write your code here
	attendace := int32(0)
	for _, t := range a {
		if t <= 0 {
			attendace++
		}
	}

	if attendace >= k {
		return "NO"
	}

	return "YES"
}

func main() {
	fmt.Println(angryProfessor(3, []int32{-1, -3, 4, 2}))
	fmt.Println(angryProfessor(2, []int32{0, -1, 2, 1}))
}
