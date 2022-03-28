package main

import "fmt"

func main() {
	var str string = ""
	for {
		str += "A"
		fmt.Println(str)
		if 10 < len(str) {
			str = ""
		}
	}
}
