package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := x
	thr := 0.0000000000001
	for math.Abs(z*z-x) > thr {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}
	return z
}

func e26() {
	fmt.Println(Sqrt(26))
}
