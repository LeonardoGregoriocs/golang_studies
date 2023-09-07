package main

import "fmt"

func main() {
	total := func() int {
		return sum(1, 3, 45, 53, 90) * 2
	}()

	fmt.Println(total)
}

func sum(numeros ...int) int {
	total := 0

	for _, numero := range numeros {
		total += numero
	}

	fmt.Println(total)

	return total
}
