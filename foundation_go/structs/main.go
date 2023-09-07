package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Pessoa interface {
	Ativar()
}

type Empresa struct {
	Nome string
}

func (e *Empresa) Ativar() {
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func (c *Cliente) Ativar() {
	c.Ativo = true
	fmt.Printf("O cliente %s foi ativado!\n", c.Nome)
}

func Ativacao(pessoa Pessoa) {
	pessoa.Ativar()
}

func main() {
	cliente1 := Cliente{
		Nome:  "Leonardo",
		Idade: 26,
		Ativo: false,
	}

	fmt.Printf("Nome: %s - Idade: %d - Ativo: %t\n", cliente1.Nome, cliente1.Idade, cliente1.Ativo)

	// cliente1.Ativar()

	Ativacao(&cliente1)

	fmt.Printf("Nome: %s - Idade: %d - Ativo: %t\n", cliente1.Nome, cliente1.Idade, cliente1.Ativo)

	cliente1.Cidade = "São Paulo"
	cliente1.Estado = "SP"

	fmt.Println(cliente1)

	minhaEmpresa := Empresa{
		Nome: "",
	}

	Ativacao(&minhaEmpresa)
}
