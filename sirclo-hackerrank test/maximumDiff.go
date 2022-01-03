package main

import "fmt"

func maxDifference(px []int32) int32 {
	// Write your code here
	diff := int32(-1)
	min := px[0]

	for i := 0; i < len(px); i++ {
		if px[i] > min && diff < px[i]-min {
			diff = px[i] - min
		}
		if px[i] < min {
			min = px[i]
		}
	}
	return diff

}

func main() {
	fmt.Println(maxDifference([]int32{7, 1, 2, 5}))
}
