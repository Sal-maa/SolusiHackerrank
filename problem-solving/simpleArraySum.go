package main

import "fmt"

func simpleArraySum(ar []int32) int32 {
	// Write your code here
	var input int32

	for _, v := range ar {
		input += v
	}
	return input
}
func main() {
	fmt.Println(simpleArraySum([]int32{1, 2, 3, 4, 10, 11}))
}
