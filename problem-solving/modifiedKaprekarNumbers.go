package main

import (
	"fmt"
	"strconv"
)

func kaprekarNumbers(p int32, q int32) {
	// Write your code here
	var result []int64
	for i := int64(p); i <= int64(q); i++ {

		sqr := i * i
		sqr_str := strconv.Itoa(int(sqr))
		if len(sqr_str)%2 != 0 {
			sqr_str = "0" + sqr_str
		}

		top, _ := strconv.Atoi(sqr_str[:(len(sqr_str) / 2)])
		bottom, _ := strconv.Atoi(sqr_str[(len(sqr_str) / 2):])

		if int64(top+bottom) == i {
			result = append(result, i)
		}
	}

	var res string
	for _, n := range result {
		res = res + strconv.Itoa(int(n)) + " "
	}

	if len(result) == 0 {
		res = "INVALID RANGE"
	}
	fmt.Println(res)
}

func main() {
	kaprekarNumbers(1, 100)
	kaprekarNumbers(400, 700)
	kaprekarNumbers(1, 99999)
	kaprekarNumbers(100, 300)
	kaprekarNumbers(900, 1000)
	kaprekarNumbers(1000, 10000)
	kaprekarNumbers(1, 10000)

}
