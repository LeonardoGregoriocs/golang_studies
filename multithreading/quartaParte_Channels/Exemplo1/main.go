package main

import "fmt"

// Thread 1 (Principal)
func main() {
	canal := make(chan string) // canal vazio

	// Thread 2
	go func() {
		canal <- "Olá Mundo!" // canal cheio
	}()

	msg := <-canal // canal esvazia
	fmt.Println(msg)
}
