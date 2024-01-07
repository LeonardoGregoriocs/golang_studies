package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	numeros := []int{1, 2, 3, 4, 5}
	for k, v := range numeros {
		fmt.Println(k, v)
	}

	// Parecido com while de outras linguagens
	i := 0
	for i < 10 {
		fmt.Println(i)
	}

	// Loop Infinito
	for {
		fmt.Println("Hello")
	}

}
