package main

import "fmt"

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	// Handle negative exponent
	if n < 0 {
		x = 1 / x
		n = -n
	}

	result := 1.0
	currentProduct := x

	for n > 0 {
		if n%2 == 1 {
			result *= currentProduct
		}
		currentProduct *= currentProduct
		n /= 2
	}

	return result
}

func main() {
	result := myPow(0.44528, 0)
	fmt.Println(result)
}
