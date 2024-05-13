package main

import "fmt"

func RevArray[T any](arr []T) []T {
	var revArray = make([]T, len(arr))
	index_to_copy := len(arr) - 1
	for _, elem := range arr {
		revArray[index_to_copy] = elem
		index_to_copy--
	}

	return revArray
}

func main() {

	fmt.Println("Reverse Array in place")
	reversed_array := RevArray([]int{1, 2, 3, 4, 5})
	fmt.Println(reversed_array)

	reversed_strings := RevArray([]string{"hello", "world", "how", "are", "you"})
	fmt.Println(reversed_strings)
}
