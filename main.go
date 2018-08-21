package main

import (
	"fmt"
	"sample/point"
)

func main() {
	p := point.NewPoint()
	fmt.Println(p.GetX())
	point.Increment(p)
	fmt.Println(p.GetX())
	p.Increment()
	fmt.Println(p.GetX())
}
