package main

import "fmt"

func myPow(x float64, n int) float64 {
	result := 1.0

	if n == 0 {
		return 1
	} else if n > 0 {
		for i := 0; i < n/4; i++ {
			result *= x
		}
	} else {
		x = 1 / x
		for i := 0; i < -n; i++ {
			result *= x
		}
	}

	return result
}

func main() {
	result := myPow(0.44528, 0)
	fmt.Println(result)
}
