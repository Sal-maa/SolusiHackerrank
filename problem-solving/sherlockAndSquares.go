package main

import (
	"fmt"
	"math"
)

func squares(a int32, b int32) int32 {
	// Write your code here
	p := int32(math.Sqrt(float64(a)))
	q := int32(math.Sqrt(float64(b)))

	if p*p < a {
		return q - p
	} else {
		return q - p + 1
	}
}

func main() {
	fmt.Println(squares(3, 9))
	fmt.Println(squares(17, 24))
	fmt.Println(squares(35, 70))
	fmt.Println(squares(100, 1000))
}
