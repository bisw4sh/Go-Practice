package main

import (
	"fmt"
	"strings"
)

func lengthOfLongestSubstring(s string) int {

	long_substring := ""
	// last_char := ""
	var seq_set []string

	for _, character := range s {

		exists := strings.Contains(long_substring, string(character)) //character is contained in the sequence string

		if exists {

			lastIndex := strings.LastIndex(long_substring, string(character))

			long_substring = long_substring[lastIndex+1:]

			long_substring = long_substring + string(character)
			// ^ This is the problem I need to move it somewhere, first extrpolate the string from the lastIndex of the character

		} else {
			long_substring += string(character)
			seq_set = append(seq_set, long_substring)
		}
		// last_char = string(character)
		fmt.Println(long_substring)
	}

	for _, elem := range seq_set {
		// fmt.Println(elem)
		if len(elem) > len(long_substring) {
			long_substring = elem
		}
	}

	return len(long_substring)
}

func main() {
	result1 := lengthOfLongestSubstring("abcabcbb")
	result2 := lengthOfLongestSubstring("bbbbb")
	result3 := lengthOfLongestSubstring("pwwkew")
	result4 := lengthOfLongestSubstring("dvdf")

	fmt.Print(result1)
	fmt.Print(result2)
	fmt.Print(result3)
	fmt.Print(result4)

}
