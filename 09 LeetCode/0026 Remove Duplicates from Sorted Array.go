package main

import "fmt"

func removeDuplicates(nums []int) int {
	//edge case of 1 length
	if len(nums) == 1 {
		return 1
	}

	unique_index := 1
	last_elem := nums[0]

	for i := 1; i < len(nums); i++ {
		if last_elem == nums[i] {
			continue
		}

		nums[unique_index] = nums[i]
		last_elem = nums[i]
		unique_index++

	}

	return unique_index
}

func main() {
	dup_array := []int{1, 2, 2, 2, 2, 3, 4, 5, 6, 7, 7, 8}
	result := removeDuplicates(dup_array)
	fmt.Println(dup_array)
	fmt.Printf("No of unique elements %d", result)
}
