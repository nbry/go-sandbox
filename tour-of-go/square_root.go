package tourofgo

import (
	"fmt"
)

const precision float64 = 0.0000000000000001

func isClose(x, y float64) bool {
	if y > x {
		return y-precision < x
	}

	return y+precision > x
}

// Get the square root of a number
func SquareRoot(x float64) float64 {
	z := 1.0

	for i := 0; i < 1000000; i++ {
		z -= (z*z - x) / (2 * z)

		if isClose(x, z*z) {
			fmt.Println("i: ", i)
			return z
		}
	}

	return z
}
