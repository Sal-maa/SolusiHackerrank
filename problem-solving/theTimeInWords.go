package main

import "fmt"

func timeInWords(h int32, m int32) string {
	// Write your code here
	hours := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve"}
	minutes := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve", "thirteen", "fourteen", "quarter", "sixteen", "seventeen", "eighteen", "nineteen", "twenty", "twenty one", "twenty two", "twenty three", "twenty four", "twenty five", "twenty six", "twenty seven", "twenty eight", "twenty nine", "half"}

	if m == 0 {
		return hours[h] + " o' clock"
	} else if m > 0 && m <= 30 {
		if m == 15 || m == 30 {
			return minutes[m] + " past " + hours[h]
		} else if m == 1 {
			return minutes[m] + " minute past " + hours[h]
		}
		return minutes[m] + " minutes past " + hours[h]
	} else if m > 30 {
		d := 60 - m
		if d == 15 {
			return minutes[d] + " to " + hours[h+1]
		} else if d == 1 {
			return minutes[d] + " minute to " + hours[h+1]
		}
		return minutes[d] + " minutes to " + hours[h+1]
	}

	return ""
}

func main() {
	fmt.Println(timeInWords(5, 47))
	fmt.Println(timeInWords(3, 0))
	fmt.Println(timeInWords(7, 29))
	fmt.Println(timeInWords(5, 30))
	fmt.Println(timeInWords(5, 45))
	fmt.Println(timeInWords(4, 15))
	fmt.Println(timeInWords(7, 30))
	fmt.Println(timeInWords(7, 15))
}
