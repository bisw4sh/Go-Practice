package main

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	var sequence, longest_sequence string
	short_ref_idx := 0
	var sequence_slice []string

	//taking shortest string for the reference to make less iterations
	for i := 1; i < len(strs); i++ {
		if len(strs[i]) < len(strs[short_ref_idx]) {
			short_ref_idx = i
		}
	}

	for _, character_of_ref := range strs[short_ref_idx] {
		sequence = sequence + string(character_of_ref)

		//loop over all the words to check if it exists there
		for _, word := range strs {
			exists := strings.Contains(word, sequence)
			if !exists {
				sequence = ""
				break
			}
		}
		sequence_slice = append(sequence_slice, sequence) //append all the sequences if found
	}


	//Find the longest sequence from the slice of sequences
	for _, elem := range sequence_slice {
		if len(elem) > len(longest_sequence) {
			longest_sequence = elem
		}
	}

	return longest_sequence //return the longest sequence
}

func main() {

	str_arr_1 := []string{"flower", "flow", "flight"}
	str_arr_2 := []string{"dog", "racecar", "car"}
	str_arr_3 := []string{"reflower","flow","flight"}

	fmt.Println(longestCommonPrefix(str_arr_1))
	fmt.Println(longestCommonPrefix(str_arr_2))
	fmt.Println(longestCommonPrefix(str_arr_3))

}
