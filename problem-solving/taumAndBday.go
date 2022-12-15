package main

import "fmt"

func taumBday(b int32, w int32, bc int32, wc int32, z int32) int64 {
	// Write your code here
	if bc > wc+z {
		return int64(b)*(int64(wc)+int64(z)) + int64(w)*int64(wc)
	} else if wc > bc+z {
		return int64(w)*(int64(bc)+int64(z)) + int64(b)*int64(bc)
	} else {
		return int64(w)*int64(wc) + int64(b)*int64(bc)
	}
}

func main() {
	fmt.Println(taumBday(27984, 1402, 619246, 615589, 247954))
	fmt.Println(taumBday(100, 100, 10000, 10000, 0))
	fmt.Println(taumBday(443463982, 833847400, 429734889, 628664883, 610875522))
	fmt.Println(taumBday(742407782, 90529439, 847666641, 8651519, 821801924))
	fmt.Println(taumBday(33588939, 88855554, 843075, 218363, 369538))
}
