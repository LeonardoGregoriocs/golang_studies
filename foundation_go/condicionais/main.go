package main

import "fmt"

func main() {
	a := 1
	b := 2
	c := 3

	if a > b && b < c {
		fmt.Println(c)
	}

	switch a {
	case 1:
		fmt.Println(a)
	case 2:
		fmt.Println(b)
	case 3:
		fmt.Println(c)
	default:
		fmt.Println("Erro")
	}
}
