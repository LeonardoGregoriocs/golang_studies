package main

import "fmt"

type Conta struct {
	saldo int
}

func (c Conta) simular(valor int) int {
	c.saldo += valor
	fmt.Printf("Simulação: %d\n", c.saldo)
	return c.saldo
}

func (c *Conta) emprestimo(valor int) int {
	c.saldo += valor
	fmt.Println("Emprestimo realizado!")
	return c.saldo
}

func main() {
	cliente1 := Conta{
		saldo: 100,
	}

	cliente1.simular(200)
	fmt.Printf("Saldo Atual após simulação: %d\n", cliente1.saldo)

	cliente1.emprestimo(200)
	fmt.Printf("Saldo Atual após emprestimo: %d\n", cliente1.saldo)
}
