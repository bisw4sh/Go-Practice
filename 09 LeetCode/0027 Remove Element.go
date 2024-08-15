package main

import (
	"fmt"
)

func removeElement(nums []int, val int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums = append(nums[:(i)], nums[(i+1):]...)
			i--
		}
	}
	return len(nums)
}

func main() {

	array_1 := []int{3, 2, 2, 3}
	array_2 := []int{0, 1, 2, 2, 3, 0, 4, 2}
	array_3 := []int{1}
	array_4 := []int{2}
	array_5 := []int{3, 3}
	array_6 := []int{4, 5}

	result_1 := removeElement(array_1, 3)
	result_2 := removeElement(array_2, 2)
	result_3 := removeElement(array_3, 1)
	result_4 := removeElement(array_4, 3)
	result_5 := removeElement(array_5, 5)
	result_6 := removeElement(array_6, 5)

	fmt.Printf("\nExpected 2 Output -> %d\n", result_1)
	fmt.Printf("\nExpected 5 Output -> %d\n", result_2)
	fmt.Printf("\nExpected 0 Output -> %d\n", result_3)
	fmt.Printf("\nExpected 1 Output -> %d\n", result_4)
	fmt.Printf("\nExpected 2 Output -> %d\n", result_5)
	fmt.Printf("\nExpected 1 Output -> %d\n", result_6)
}
