package main

import (
	"fmt"
)

func mySqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}

	left, right := 1, x
	for left <= right {
		mid := left + (right-left)/2
		if mid*mid == x {
			return mid
		} else if mid*mid < x {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return right
}

func main() {
	result_1 := mySqrt(4)
	fmt.Println(result_1)

	result_2 := mySqrt(8)
	fmt.Println(result_2)

	result_3 := mySqrt(150)
	fmt.Println(result_3)

	result_4 := mySqrt(170)
	fmt.Println(result_4)
}
