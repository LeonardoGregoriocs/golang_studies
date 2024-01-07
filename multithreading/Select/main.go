package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go func() {
		time.Sleep(time.Second * 4)
		c1 <- 1
	}()
	go func() {
		time.Sleep(time.Second * 4)
		c2 <- 2
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		case <-time.After(time.Second * 3):
			fmt.Println("timeout")
		default:
			fmt.Println("default")
		}
	}

	// Para recebermos mensagem sem parar, podemos fazer um for sem condição:
	for {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		case <-time.After(time.Second * 3):
			fmt.Println("timeout")
		default:
			fmt.Println("default")
		}
	}

}
