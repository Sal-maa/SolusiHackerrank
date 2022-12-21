package main

import "fmt"

func contains(s []int32, str int32) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func beautifulTriplets(d int32, arr []int32) int32 {
	// Write your code here
	check := make(map[int32]int32)
	for i, k := range arr {
		if contains(arr, k+d) {
			check[int32(i)] = k + d
		}
	}

	count := int32(0)
	for q := range check {
		for _, v := range check {
			if arr[q] == v {
				count++
			}
		}
	}
	return count
}

func main() {
	fmt.Println(beautifulTriplets(3, []int32{1, 6, 7, 7, 8, 10, 12, 13, 14, 19}))
	fmt.Println(beautifulTriplets(10, []int32{1, 2, 3, 4, 5}))
	fmt.Println(beautifulTriplets(1, []int32{1, 1}))
	fmt.Println(beautifulTriplets(2, []int32{2, 4}))
	fmt.Println(beautifulTriplets(3, []int32{1, 2, 4, 5, 7, 8, 10}))
}
