package main

import "fmt"

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	} else if n == 1 {
		return x
	} else if n < 0 {
		return 1 / myPow(x, -n)
	}

	return x * myPow(x, n-1)
}

func main() {
	result := myPow(0.44528, 0)
	fmt.Println(result)
}
