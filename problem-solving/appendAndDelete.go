package main

import (
	"fmt"
	"strings"
)

func appendAndDelete(s string, t string, k int32) string {
	// Write your code here
	sSlice := strings.Split(s, "")
	tSlice := strings.Split(t, "")

	mapS := make(map[int]string, 0)
	for idx, str := range tSlice {
		mapS[idx] = string(str)
	}

	var init int = 0
	for idx, str := range sSlice {
		if mapS[idx] != string(str) {
			init = idx
			break
		}
	}

	step1 := len(sSlice[init:])
	step2 := len(tSlice[init:])

	var c bool
	if len(tSlice) > len(sSlice) {
		c = s == string(t[len(tSlice)-len(sSlice):])
	}

	if int(k) >= (step1+step2) || s == t || c {
		return "Yes"
	}

	return "No"
}

func main() {
	fmt.Println(appendAndDelete("hackerhappy", "hackerrank", 9)) //yes
	fmt.Println(appendAndDelete("aba", "aba", 7))                //yes
	fmt.Println(appendAndDelete("y", "yu", 2))                   //no
}
