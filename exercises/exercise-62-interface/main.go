package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
	width  float64
}

type circle struct {
	radius float64
}

func (s square) Area() float64 {
	return s.length * s.width
}

func (c circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

type shape interface {
	Area() float64
}

func Info(sh shape) {
	fmt.Println(sh.Area())
}

func main() {
	sq := square{
		length: 5,
		width:  8,
	}

	cir := circle{
		radius: 4,
	}
	Info(sq)
	Info(cir)
}
