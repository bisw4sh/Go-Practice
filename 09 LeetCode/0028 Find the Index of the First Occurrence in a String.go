package main

import (
	"fmt"
	"strings"
)

func strStr(haystack string, needle string) int {

	if !strings.Contains(haystack, needle) {
		return -1
	}

	return strings.Index(haystack, needle)
}

func main() {
	result_1 := strStr("sadbutsad", "sad")
	fmt.Println(result_1)

	result_2 := strStr("leetcode", "leeto")
	fmt.Println(result_2)

	result_3 := strStr("mississippi", "issip")
	fmt.Println(result_3)
}
