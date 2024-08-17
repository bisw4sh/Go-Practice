package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	s = strings.Trim(s, " ")
	slice_strings := strings.Split(s, " ")

	return len(slice_strings[len(slice_strings)-1])
}

func main() {
	result_1 := lengthOfLastWord("Hello World")
	fmt.Println(result_1)

	result_2 := lengthOfLastWord("   fly me   to   the moon  ")
	fmt.Println(result_2)

	result_3 := lengthOfLastWord("luffy is still joyboy")
	fmt.Println(result_3)
}
