package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 1.0

	fmt.Println("-----")
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}

	return z
}

func main() {
	nums := []float64{2, 3, 4, 5}
	for _, num := range nums {
		fmt.Println(Sqrt(num))
	}
}
