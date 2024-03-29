package main

import "fmt"

func libraryFine(d1 int32, m1 int32, y1 int32, d2 int32, m2 int32, y2 int32) int32 {
	// Write your code here
	var fine int32 = 0
	if d1-d2 > 0 && m1-m2 == 0 && y1-y2 == 0 {
		fine = 15 * (d1 - d2)
	} else if m1-m2 > 0 && y1-y2 == 0 {
		fine = 500 * (m1 - m2)
	} else if y1-y2 > 0 {
		fine = 10000
	}

	return fine
}

func main() {
	fmt.Println(libraryFine(6, 6, 2015, 9, 6, 2015))
	fmt.Println(libraryFine(30, 5, 2015, 2, 5, 2015))
	fmt.Println(libraryFine(2, 7, 1014, 1, 1, 1014))
	fmt.Println(libraryFine(2, 7, 2014, 1, 2, 2015))
}
