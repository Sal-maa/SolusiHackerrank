package main

import (
	"fmt"
	"math/big"
)

func extraLongFactorials(n int32) {
	// Write your code here
	var fact big.Int
	fact.MulRange(1, int64(n))
	num := fact.String()
	fmt.Println(num)
}

func main() {
	extraLongFactorials(25)
	extraLongFactorials(37)
	extraLongFactorials(58)
}
