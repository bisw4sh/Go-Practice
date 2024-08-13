package main

import "fmt"

func twoSum(nums []int, target int) []int {

	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	res := twoSum(arr, 9)
	fmt.Println(res)
}