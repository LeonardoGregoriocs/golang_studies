package main

import "github.com/leonardogregoriocs/golang_studies/manipulando_eventos/pkg/rabbitmq"

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	rabbitmq.Publish(ch, "Hello, world! ")
}
