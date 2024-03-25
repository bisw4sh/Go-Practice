package main

import "fmt"

func checking() {
	fmt.Println("From out the main function")
}

func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
	checking()
	checking()
	checking()
	checking()

	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
