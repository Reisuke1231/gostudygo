package main

import (
	"fmt"
	"time"
)

func main() {
	// t := time.Now()
	t := time.Date(2022, 3, 28, 21, 0, 0, 0, time.Local)
	fmt.Println(t)
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	case t.Hour() > 21:
		fmt.Println("Good night.")
	default:
		fmt.Println("Good evening.")
	}
}
