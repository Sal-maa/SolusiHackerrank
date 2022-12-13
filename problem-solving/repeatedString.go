package main

import (
	"fmt"
	"strings"
)

func repeatedString(s string, n int64) int64 {
	// Write your code here
	if !strings.Contains(s, "a") {
		return 0
	}

	if n > int64(len(s)) {
		sisa := n % int64(len(s))
		counter := n / int64(len(s))

		return int64(strings.Count(s, "a"))*counter + int64(strings.Count(s[:sisa], "a"))
	} else {
		return int64(strings.Count(s[:n], "a"))
	}
}

func main() {
	fmt.Println(repeatedString("ababa", 3))
	fmt.Println(repeatedString("d", 590826798023))
	fmt.Println(repeatedString("beeaabc", 711560125001))
	fmt.Println(repeatedString("ebcdddafdfeffaddbceddebafbbebebbbcefcbcdfbaabecfaaeeaaffdfccffbdeeaabcfeecfcecbfebacefebdfaeedadebdf", 561984209086))
	fmt.Println(repeatedString("aadcdaccacabdaabadadaabacdcbcacabbbadbdadacbdadaccbbadbdcadbdcacacbcacddbcbbbaaccbaddcabaacbcaabbaaa", 942885108885))
	fmt.Println(repeatedString("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", 534802106762))
	fmt.Println(repeatedString("ojowrdcpavatfacuunxycyrmpbkvaxyrsgquwehhurnicgicmrpmgegftjszgvsgqavcrvdtsxlkxjpqtlnkjuyraknwxmnthfpt", 685118368975))
	fmt.Println(repeatedString("udjlitpopjhipmwgvggazhuzvcmzhulowmveqyktlakdufzcefrxufssqdslyfuiahtzjjdeaxqeiarcjpponoclynbtraaawrps", 872514961806))
	fmt.Println(repeatedString("aba", 10))
	fmt.Println(repeatedString("aab", 882787))
	fmt.Println(repeatedString("gfcaaaecbg", 547602))
	fmt.Println(repeatedString("jdiacikk", 899491))
}
