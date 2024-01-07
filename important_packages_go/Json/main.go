package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Conta struct {
	Numero int `json:"n"`
	Saldo  int `json:"s"`
}

func main() {
	conta := Conta{
		Numero: 1,
		Saldo:  100,
	}

	res, err := json.Marshal(conta)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(res))

	enconder := json.NewEncoder(os.Stdout)
	enconder.Encode(conta)

	var contaX Conta

	err = json.Unmarshal(res, &contaX)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(contaX)
}
