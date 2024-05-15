package main

import (
	"errors"
	"fmt"
)

func prod(arr []int) (int, error) {
	mul := 1

	for _, elem := range arr {

		if elem == 0 {
			return 0, errors.New("a zero element")
		}
		mul *= elem
	}

	return mul, nil
}

func main() {
	no_zero, err1 := prod([]int{1, 2, 3, 4, 5, 5})
	with_zero, err2 := prod([]int{1, 2, 3, 4, 5, 0})

	if err1 == nil {
		fmt.Println(no_zero)
	} else {
		fmt.Println(err1)
	}

	if err2 == nil {

		fmt.Println(with_zero)
	} else {
		fmt.Println(err2)
	}

}
