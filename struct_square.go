package main

import (
	"fmt"
)

type Point struct {
	X int
	Y int
}

func (p *Point) Move(dx int, dy int) {
	p.X += dx
	p.Y += dy
}

func (s *Square) Area() int {
	return s.Length * s.Length
}

type Square struct {
	Center Point
	Length int
}

func main() {
	p := Point{2, 3}
	fmt.Printf("%v\n", p)
	p.Move(3, 3)
	fmt.Printf("%v\n", p)

	s := Square{p, 10}
	fmt.Printf("%v\n", s)
	area := s.Area()
	fmt.Printf("%v\n", area)

}
