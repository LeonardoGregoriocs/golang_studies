package main

import (
	"github.com/leonardogregoriocs/golang_studies/apis/configs"
)

func main() {
	_, _ = configs.LoadConfig(".")
}
