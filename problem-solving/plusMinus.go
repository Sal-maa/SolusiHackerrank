package main

import "fmt"

func plusMinus(arr []int32) {
	// Write your code here
	var plus, zero int32

	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			zero++
		} else if arr[i] > 0 {
			plus++
		}
	}
	fmt.Printf("%.6f\n", float64(plus)/float64(len(arr)))
	fmt.Printf("%.6f\n", float64(int32(len(arr))-zero-plus)/float64(len(arr)))
	fmt.Printf("%.6f\n", float64(zero)/float64(len(arr)))

}

func main() {
	plusMinus([]int32{-4, 3, -9, 0, 4, 1})
}
