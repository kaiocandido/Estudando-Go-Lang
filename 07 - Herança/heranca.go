package main

import "fmt"

type Pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
}

type Estudante struct {
	Pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Herança")

	pessoa1 := Pessoa{"João", "Silva", 30}
	fmt.Println(pessoa1)

	estudante1 := Estudante{pessoa1, "Engenharia", "UFSC"}
	fmt.Println(estudante1)
	fmt.Println(estudante1.nome) // Acessando o campo nome da struct Pessoa através da struct Estudante
}
