package main

import (
	"fmt"
)

func sorting(nums []int) {
	quicksort(nums, 0, len(nums)-1)
}

func quicksort(nums []int, start, end int) {
	if start >= end {
		return
	}
	pivotIndex := partition(nums, start, end)

	quicksort(nums, start, pivotIndex-1)
	quicksort(nums, pivotIndex+1, end)
}

func partition(nums []int, start, end int) int {
	pivotValue := nums[start]
	left := start + 1
	right := end

	for left <= right {

		for left <= right && nums[left] <= pivotValue {
			left++
		}
		for left <= right && nums[right] >= pivotValue {
			right--
		}

		if left < right {
			nums[left], nums[right] = nums[right], nums[left]
		}
	}

	nums[start], nums[right] = nums[right], nums[start]

	return right
}

func main() {
	arr := []int{1, 3, 2, 4, 6, 5, 9, 8, 7, 11, 10, 12}
	sorting(arr)
	fmt.Println(arr)
}
