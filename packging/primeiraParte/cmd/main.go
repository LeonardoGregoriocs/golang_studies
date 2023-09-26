package main

import (
	"fmt"

	"github.com/leonardogregoriocs/golang_studies/packging/primeiraParte/math"
)

func main() {
	m := math.NewMath(1, 3)
	fmt.Println(m.Add())
	fmt.Println("Hello, world")
}
