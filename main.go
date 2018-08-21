package main

import (
	"fmt"

	"github.com/bols-blue-org/cgo-swig/point"
)

func main() {
	p := point.NewPoint()
	fmt.Println(p.GetX())
	point.Increment(p)
	fmt.Println(p.GetX())
	p.Increment()
	fmt.Println(p.GetX())
}
