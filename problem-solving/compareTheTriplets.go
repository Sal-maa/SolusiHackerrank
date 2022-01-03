package main

import "fmt"

func compareTriplets(a []int32, b []int32) []int32 {
	// Write your code here
	var input []int32
	input = append(input, 0)
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] < b[i] {
			input[1]++
		} else if a[i] > b[i] {
			input[0]++
		}
	}
	return input
}

func main() {
	fmt.Println(compareTriplets([]int32{17, 28, 30}, []int32{99, 16, 8}))
}
