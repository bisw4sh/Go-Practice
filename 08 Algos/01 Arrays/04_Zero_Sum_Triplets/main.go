package main

import (
	"errors"
	"fmt"
)

func triples(arr []int) (int, int, int, error) {
	for i, elemi := range arr {
		for j, elemj := range arr {
			for k, elemk := range arr {
				sum := elemi + elemj + elemk
				if sum == 0 {
					return i, j, k, nil
				}
			}
		}
	}

	return -1, -1, -1, errors.New("no triple sum equals zero")
}

func main() {
	tripleSumArr := []int{5, 10, 3, 6, -9, -15}
	i, j, k, err := triples(tripleSumArr)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(tripleSumArr[i], tripleSumArr[j], tripleSumArr[k])
	}
}
