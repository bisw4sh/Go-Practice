package main

import "fmt"

type void struct{}

func find_duplicate(arr []string) map[string]void {

	set := make(map[string]void)
	var member void

	for idx, elem := range arr {
		for i := 0; i < len(arr); i++ {
			if i == idx {
				continue
			} else if elem == arr[i] {

				set[elem] = member

			}
		}
	}

	return set
}

func main() {

	returned_duplicate_set := find_duplicate([]string{"go", "pearl", "rust", "js", "ts", "c", "c++", "go", "ts"})

	fmt.Println(returned_duplicate_set)
}
