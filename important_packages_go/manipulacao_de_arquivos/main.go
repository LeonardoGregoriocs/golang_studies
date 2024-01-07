package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Criação de arquivo
	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}

	// Gravar algo no arquivo
	// tamanho, err := f.Write([]bytes("Hello, World"))
	tamanho, err := f.WriteString("Hello, World")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Tamanho: %d\n", tamanho)
	f.Close()

	// Ler um arquivo
	arquivo, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Tamanho: %v\n", string(arquivo))

	// Leitura de pouco em pouco abrindo o arquivo

	//Abertura do arquivo
	file, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}

	// Reader --> Responsavel por ler o arquivo
	reader := bufio.NewReader(file)

	//Tamanho do buffer
	buffer := make([]byte, 10)

	// Responsavel por pegar os dados e colocar no buffer, qnd acabar, temos um break.
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}

	// Remove arquivos

	err = os.Remove("arquivo.txt")
	if err != nil {
		panic(err)
	}

}
