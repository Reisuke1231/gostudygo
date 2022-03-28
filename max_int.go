package main

import (
	"fmt"
)

func main() {
	nums := []int{16, 8, 42, 4, 23, 15}

	var max int = nums[0]
	for _, num := range nums {
		if max < num {
			max = num
		}
	}

	fmt.Println(max)
}
