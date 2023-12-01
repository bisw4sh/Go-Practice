package main

import (
	"fmt")

func main() {
	for item, index := range []int{1,2,3,4,5,6,7,8,9,10,11}{
		fmt.Printf("Item %d of index %d\n", item, index)
	}

	// mapped := map[string]int {"foo" : 1, "bar" : 2, "baz" : 3}

	for key, value := range map[string]int {"foo" : 1, "bar" : 2, "baz" : 3}{
		fmt.Printf("%s -> %d\n", key, value)
	}

}