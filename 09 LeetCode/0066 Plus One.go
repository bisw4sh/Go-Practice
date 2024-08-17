package main

import "fmt"

func plusOne(digits []int) []int {
	length := len(digits)

	if length < 1 {
		return []int{1}
	}
	carry := 0

	for i := length - 1; i >= 0; i-- {

		if i == length-1 {
			carry = digits[i] + 1
			digits[i] = carry % 10
			carry = carry / 10
		} else {
			elem_temp := digits[i]
			digits[i] = (elem_temp + carry) % 10
			carry = (elem_temp + carry) / 10
		}
	}

	if carry > 0 {
		carry_slice := []int{carry}
		digits = append(carry_slice, digits...)
	}

	return digits
}

func main() {
	result_1 := plusOne([]int{1, 2, 3})
	fmt.Println(result_1)

	result_2 := plusOne([]int{4, 3, 2, 1})
	fmt.Println(result_2)

	result_3 := plusOne([]int{9})
	fmt.Println(result_3)

	result_4 := plusOne([]int{})
	fmt.Println(result_4)

	result_5 := plusOne([]int{9, 9})
	fmt.Println(result_5)
}
