package main

import (
	"fmt"
	"strings"
)

func firstOccurrence(s string, x string) int32 {
	// Write your code here
	hasil := 0
	xSplit := strings.Split(x, "*")
	indexAwal := strings.Index(s, xSplit[0])

	if strings.Contains(s, xSplit[0]) && strings.Contains(s[indexAwal:indexAwal+len(x)], xSplit[len(xSplit)-1]) {
		hasil = indexAwal
	} /*else {
		hasil = strings.Index(s, x)
	}*/

	return int32(hasil)
}

func main() {
	fmt.Println(firstOccurrence("juliasamanthantjulia", "ant"))
	fmt.Println(firstOccurrence("juliasamanthasamanthajulia", "has"))
	fmt.Println(firstOccurrence("juliasamanthasamanthajulia", "ant*as"))
}
