package main

import "fmt"

func maxDistance(arrays [][]int) int {

	// Initialize with the first array's min and max
	min, max := arrays[0][0], arrays[0][len(arrays[0])-1]
	result := 0

	for i := 1; i < len(arrays); i++ {
		array := arrays[i]
		// Calculate potential maximum distance using current array's first and last elements
		result = maxInt(result, maxDistanceBetween(min, max, array))
		// Update global min and max
		min = minInt(min, array[0])
		max = maxInt(max, array[len(array)-1])
	}

	return result
}

// Helper function to calculate the max distance between global min, max, and the current array's elements
func maxDistanceBetween(min, max int, array []int) int {
	return maxInt(absDistance(array[len(array)-1], min), absDistance(max, array[0]))
}

// Helper function to calculate the absolute distance between two integers
func absDistance(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

// Helper functions to get the minimum and maximum of two integers
func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {

	distance_list_1 := [][]int{{1, 2, 3}, {4, 5}, {1, 2, 3}}
	max_distance_1 := maxDistance(distance_list_1)
	fmt.Println(max_distance_1)

	distance_list_2 := [][]int{{1}, {1}}
	max_distance_2 := maxDistance(distance_list_2)
	fmt.Println(max_distance_2)

	distance_list_3 := [][]int{{1, 4}, {0, 5}}
	max_distance_3 := maxDistance(distance_list_3)
	fmt.Println(max_distance_3)

	distance_list_4 := [][]int{{1}, {2}}
	max_distance_4 := maxDistance(distance_list_4)
	fmt.Println(max_distance_4)

	distance_list_5 := [][]int{{1, 4}, {0, 5}}
	max_distance_5 := maxDistance(distance_list_5)
	fmt.Println(max_distance_5)

}
