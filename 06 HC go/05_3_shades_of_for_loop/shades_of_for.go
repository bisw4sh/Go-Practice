package main

import "fmt"

func main() {
	sum := 0
//all 3 seperations
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
//all 3 seperations with semicolon
	summ := 0

		for ; summ < 1000; {
		summ += summ
	}
	fmt.Println(summ)

	//ommited semicolons too
	summm := 0

		for summm < 1000 {
		sum += summm
	}
	fmt.Println(summm)


	// Basically a While True
	// 	for {

	// }
}