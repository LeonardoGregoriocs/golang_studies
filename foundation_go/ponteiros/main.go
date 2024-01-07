package main

import "fmt"

func main() {
	a := 10

	// Endereço na memória
	fmt.Println(&a)

	var b *int = &a

	// Mostra o endereço na memória da variavel b, mesmo de a.
	fmt.Println(b)

	// Mostrar o valor da variável a e b
	fmt.Println(*b)

	c := &a

	// Mostra o endereço na memória da variavel c, mesmo de a e b.
	fmt.Println(c)

	*c = 25

	fmt.Println(a)
	fmt.Println(*b)
	fmt.Println(*c)
}
