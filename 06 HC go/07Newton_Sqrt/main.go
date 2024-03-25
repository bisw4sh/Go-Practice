package main

import (
	"fmt"
)

func functionalVal(x, initial, final, temp float64) (ansinitial, ansfinal, ansTemp float64) {
	ansinitial = initial*initial - x
	ansfinal = final*final - x
	ansTemp = temp*temp - x
	return
}
func Sqrt(x float64) float64 {

	var initial, final, temp, ansInitial, ansFinal, ansTemp float64
	initial = 0.0
	final = 5.0

	for i := 0; i < 10; i++ {

		temp = (initial + final) / 2
		ansInitial, ansFinal, ansTemp = functionalVal(x, initial, final, temp)

		if ansInitial < 0 && ansFinal > 0 {
			if ansTemp < 0 {
				initial = temp
			} else if ansTemp > 0 {
				final = temp
			}
		} else if ansInitial > 0 && ansFinal < 0 {
			if ansTemp > 0 {
				initial = temp
			} else if ansTemp < 0 {
				final = temp
			}
		}

	}

	return temp
}

func main() {
	fmt.Println(Sqrt(2))
}
