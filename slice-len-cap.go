package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// panic: runtime error: slice bounds out of range [:8] with capacity 6
	// s = s[:8]
	// printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)

	// panic: runtime error: slice bounds out of range [:6] with capacity 4
	// s = s[0:6]
	// printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
