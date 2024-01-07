package main

import "fmt"

const x = "Hello, World!"

var (
	b bool    = true
	c int     = 10
	d string  = "Wesley"
	e float64 = 1.2
)

func main() {
	a := "X" // string
	println(a)
	fmt.Println(a, b, c, d, e, x)
}
