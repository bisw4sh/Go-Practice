package main

import "fmt"

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	for index, element := range nums {

		if element >= target {
			return index
		}
	}

	return len(nums)
}

func main() {
	result_1 := searchInsert([]int{1, 3, 5, 6}, 5)
	fmt.Println(result_1)

	result_2 := searchInsert([]int{1, 3, 5, 6}, 2)
	fmt.Println(result_2)

	result_3 := searchInsert([]int{1, 3, 5, 6}, 7)
	fmt.Println(result_3)
}
