package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	var sequence string
	short_ref_idx := 0

	//taking shortest string for the reference to make less iterations
	for i := 1; i < len(strs); i++ {
		if len(strs[i]) < len(strs[short_ref_idx]) {
			short_ref_idx = i
		}
	}

	for i := 0; i < len(strs[short_ref_idx]); i++ {
		for word := 0; word < len(strs); word++ {
			if strs[short_ref_idx][i] != strs[word][i] {
				return sequence
			}
		}
		sequence = sequence + string(strs[short_ref_idx][i])
	}

	return sequence
}

func main() {

	str_arr_1 := []string{"flower", "flow", "flight"}
	str_arr_2 := []string{"dog", "racecar", "car"}
	str_arr_3 := []string{"reflower", "flow", "flight"}

	fmt.Println(longestCommonPrefix(str_arr_1))
	fmt.Println(longestCommonPrefix(str_arr_2))
	fmt.Println(longestCommonPrefix(str_arr_3))

}
