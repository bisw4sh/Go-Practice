package main

import (
	"fmt"
)

func adder(list1, list2 []interface{}) []interface{} {
	shortestLength := len(list1)
	longestLength := len(list2)
	if shortestLength > longestLength {
		shortestLength = len(list2)
		longestLength = len(list1)
	}

	addedArr := make([]interface{}, longestLength)

	for i := 0; i < shortestLength; i++ {
		switch v1 := list1[i].(type) {
		case int:
			switch v2 := list2[i].(type) {
			case int:
				addedArr[i] = v1 + v2
			}
		case float64:
			switch v2 := list2[i].(type) {
			case float64:
				addedArr[i] = v1 + v2
			}
		case string:
			switch v2 := list2[i].(type) {
			case string:
				addedArr[i] = v1 + v2
			}
		}
	}

	for i := shortestLength; i < longestLength; i++ {
		if longestLength == len(list1) {
			addedArr[i] = list1[i]
		} else {
			addedArr[i] = list2[i]
		}
	}

	return addedArr
}

func main() {
	retAddedArr := adder([]interface{}{1, 2, 3, 4, 5, 5.5}, []interface{}{1, 2, 3, 4, 5})
	fmt.Println(retAddedArr)

	retAddedArr = adder([]interface{}{"hello", " ", "world"}, []interface{}{"hello", "gopher"})
	fmt.Println(retAddedArr)
}
