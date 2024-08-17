package main

import "fmt"

func singleNumber(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	set := make(map[int]bool)

	for _, elem := range nums {
		_, exists := set[elem]
		if exists {
			delete(set, elem)
		} else {
			set[elem] = true
		}
	}
	var element int

	for key := range set {
		element = key
	}

	return element
}

func main() {
	result_1 := singleNumber([]int{2, 2, 1})
	fmt.Println(result_1)

	result_2 := singleNumber([]int{4, 1, 2, 1, 2})
	fmt.Println(result_2)

	result_3 := singleNumber([]int{1})
	fmt.Println(result_3)
}
