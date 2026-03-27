package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0

	// for para repetir um bloco de código enquanto uma condição for verdadeira
	for i < 5 {
		i++
		println(i)
		time.Sleep(time.Second)

	}

	fmt.Println(i)

	// for para repetir um bloco de código um número específico de vezes
	for x := 0; x < 10; x++ {
		fmt.Println(x)
		time.Sleep(time.Second)
	}

	nome := []string{"João", "Maria", "Pedro"}

	// for para iterar sobre uma coleção de dados
	for _, nome := range nome {
		fmt.Println("Nome", nome)
	}

	// for para iterar sobre uma string
	for _, letra := range "Golang" {
		fmt.Println(string(letra))
	}

	// for para iterar sobre um map
	mapa := map[string]int{
		"João":  30,
		"Maria": 25,
		"Pedro": 20,
	}

	for indice, idade := range mapa {
		fmt.Printf("Chave: %s, Valor: %d\n", indice, idade)
	}

}
