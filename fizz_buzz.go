package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str string
	for i := 1; i <= 50; i++ {
		str = ""
		str += fizz(i)
		str += buzz(i)
		if str == "" {
			str += strconv.Itoa(i)
		}
		fmt.Printf("%s\n", str)
	}
}

func fizz(num int) string {
	reply := ""
	if num%3 == 0 {
		reply = "Fizz"
	}
	return reply
}

func buzz(num int) string {
	reply := ""
	if num%5 == 0 {
		reply = "Buzz"
	}
	return reply
}
