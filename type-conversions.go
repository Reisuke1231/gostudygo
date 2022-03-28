package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)

	i := 42
	i_to_f := float64(i)
	u := uint(i_to_f)
	fmt.Printf("%T: %v\n", x, x)
	fmt.Printf("%T: %v\n", y, y)
	fmt.Printf("%T: %v\n", f, f)
	fmt.Printf("%T: %v\n", z, z)
	fmt.Printf("%T: %v\n", i, i)
	fmt.Printf("%T: %v\n", i_to_f, i_to_f)
	fmt.Printf("%T: %v\n", u, u)
}
