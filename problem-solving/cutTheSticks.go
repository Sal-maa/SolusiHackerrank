package main

import "fmt"

func countCuting(arr []int32) []int32 {
	if len(arr) == 0 {
		return []int32{0}
	}

	var newarr []int32
	min := arr[0]
	for _, number := range arr {
		if number <= min {
			min = number
		}
	}

	for _, number := range arr {
		newnumber := number - min
		if newnumber != 0 {
			newarr = append(newarr, newnumber)
		}
	}
	return newarr
}

func cutTheSticks(arr []int32) []int32 {
	// Write your code here
	var result []int32

	for i := 0; ; i++ {
		if len(arr) == 0 {
			break
		}
		result = append(result, int32(len(arr)))
		arr = countCuting(arr)
	}

	return result
}

func main() {
	fmt.Println(cutTheSticks([]int32{66, 42, 68, 72, 68, 81, 91, 19, 40, 73, 44, 73, 89, 73, 44, 12, 77, 40, 44, 17, 59, 99, 35, 92, 80, 51, 14, 30}))
	fmt.Println(cutTheSticks([]int32{8, 8, 14, 10, 3, 5, 14, 12}))
	fmt.Println(cutTheSticks([]int32{1, 13, 3, 8, 14, 9, 4, 4}))
	fmt.Println(cutTheSticks([]int32{4, 5, 10, 8, 11}))
}
