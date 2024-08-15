package main

import "fmt"

func lemonadeChange(bills []int) bool {

	// bill_type := [3]int{5, 10, 20} //only types of bills
	number_of_bills_in_index := [3]int{0, 0, 0}

	for _, bill_selected := range bills {

		switch {

		case bill_selected == 5:
			number_of_bills_in_index[0]++

		case bill_selected == 10:
			if number_of_bills_in_index[0] < 1 {
				//No $5 bill means, can't return
				return false
			} else {
				number_of_bills_in_index[1]++
				number_of_bills_in_index[0]--
			}

		case bill_selected == 20:

			if number_of_bills_in_index[0] > 0 {

				if number_of_bills_in_index[1] > 0 {
					number_of_bills_in_index[0]--
					number_of_bills_in_index[1]--
					number_of_bills_in_index[2]++
				} else if number_of_bills_in_index[0] >= 3 && number_of_bills_in_index[1] == 0 {
					number_of_bills_in_index[0] -= 3
					number_of_bills_in_index[2]++
				} else {
					return false
				}
			} else {
				//No $5 bill means, can't return
				return false
			}

		default:
			return false
		}
	}

	return true
}

func main() {
	//First test case
	bills_1 := []int{5, 5, 5, 10, 20}
	result_1 := lemonadeChange(bills_1)
	fmt.Println(result_1)

	//Second test case
	bills_2 := []int{5, 5, 10, 10, 20}
	result_2 := lemonadeChange(bills_2)
	fmt.Println(result_2)
}
