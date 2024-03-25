package main

import "fmt"

func swap( a, b string) (string, string){
	return b,a 
}

func main() {
	fmt.Println("Hello")
	a:= "apple"
	b:="banana"
	fmt.Printf("Unswapped ver. %s and %s\n", a, b)
	a,b = swap(a,b)
	fmt.Printf("swapped ver. %s and %s\n", a, b)

	x,y := split(5)
	fmt.Printf("Value of x -> %d and y -> %d\n", x,y)
}

//named returns
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}