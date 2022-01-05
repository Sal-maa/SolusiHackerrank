package main

import (
	"fmt"
	"strconv"
	"strings"
)

func timeConversion(s string) string {
	// Write your code here
	var t []string
	hConv, _ := strconv.Atoi(s[:2])

	if hConv == 12 && s[len(s)-2:] == "AM" {
		t = append(t, "00")
	} else if hConv < 12 && s[len(s)-2:] == "PM" {
		hConv = hConv + 12
		t = append(t, strconv.Itoa(hConv))
	} else {
		t = append(t, s[:2])
	}
	t = append(t, string(s[2:len(s)-2]))
	tstring := strings.Join(t, "")
	return tstring
}

func main() {
	fmt.Println(timeConversion("07:05:45PM"))
}
