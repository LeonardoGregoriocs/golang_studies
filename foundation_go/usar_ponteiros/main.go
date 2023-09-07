package main

import "fmt"

func soma(a, b *int) int {
	*a = 40
	return *a + *b
}

func main() {

	minhaVar1 := 10
	minhaVar2 := 20

	fmt.Println(soma(&minhaVar1, &minhaVar2))
}
