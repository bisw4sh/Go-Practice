package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	left_median := 0
	right_median := 0
	num1_pointer_idx := 0
	num2_pointer_idx := 0

	len_sorted := len(nums1) + len(nums2)

	var sorted_array [2000]int

	for i := 0; i <= len_sorted/2; i++ {

		if num1_pointer_idx >= len(nums1) {
			//Need no concern nums1 array, it is completed
			sorted_array[i] = nums2[num2_pointer_idx]
			num2_pointer_idx++

		} else if num2_pointer_idx >= len(nums2) {

			//Need no concern nums2 array, it is completed
			sorted_array[i] = nums1[num1_pointer_idx]
			num1_pointer_idx++

		} else if num1_pointer_idx < len(nums1) && nums1[num1_pointer_idx] <= nums2[num2_pointer_idx] {

			//len of nums1 in bound but gotta check if it is smaller or equal nums2
			sorted_array[i] = nums1[num1_pointer_idx]
			num1_pointer_idx++

		} else if num2_pointer_idx < len(nums2) && nums2[num2_pointer_idx] < nums1[num1_pointer_idx] {

			sorted_array[i] = nums2[num2_pointer_idx]
			num2_pointer_idx++

		}

		//left = n/2 -1 and right = n/2
		left_median = right_median
		right_median = sorted_array[i]

		// fmt.Println(sorted_array[i])
	}

	//Return by calc the half if the length of the array is even
	if len_sorted%2 == 0 {
		return float64(left_median+right_median) / 2.0
	}

	//For odd number of length
	return float64(right_median)
}

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}

	result := findMedianSortedArrays(nums1, nums2)
	fmt.Println(result)
}
