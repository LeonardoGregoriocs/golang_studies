package main

import (
	"fmt"

	"github.com/leonardogregoriocs/fcutils-secret/pkg/events"
)

func main() {
	ed := events.NewEventDispatcher()
	fmt.Println(ed)
}

// Como fazer com que o Go baixe informações de repositório privado?
// Alteramos a váriavel de ambiente GOPRIVATE
