package shortened_if

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func Main2() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}