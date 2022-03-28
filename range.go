package main

import (
	"fmt"
)

func main() {
	loons := []string{"bugs", "daffy", "taz"}

	fmt.Println("----")
	for i := 0; i < len(loons); i++ {
		fmt.Println(loons[i])
	}

	fmt.Println("----")
	for i := range loons {
		fmt.Println(i)
	}

	fmt.Println("----")
	for i, name := range loons {
		fmt.Printf("%d: %s\n", i, name)
	}

	fmt.Println("----")
	for _, name := range loons {
		fmt.Println(name)
	}

	fmt.Println("----")
	loons = append(loons, "elmer")
	fmt.Println(loons)
}
