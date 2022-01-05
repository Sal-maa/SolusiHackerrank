package main

import "fmt"

func staircase(n int32) {
	// Write your code here
	for i := 1; i < int(n)+1; i++ {
		for j := 1; j < int(n)+1; j++ {
			if j >= int(n)-i+1 {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func main() {
	staircase(6)
}
