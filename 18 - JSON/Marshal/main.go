package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Idade int    `json:"idade"`
	Raca  string `json:"raca"`
}

func main() {
	// usado para converter um objeto Go em JSON
	//json.Marshal()
	// usado para converter um JSON em um objeto Go
	//json.Unmarshal()
	cachorro := cachorro{
		Nome:  "Rex",
		Idade: 5,
		Raca:  "Labrador",
	}
	fmt.Println(cachorro)

	convertido, error := json.Marshal(cachorro)

	if error != nil {
		fmt.Println("Erro ao converter para JSON:", error)
	}

	// o resultado é um slice de bytes, então precisamos converter para string ou usar bytes.NewBuffer para imprimir
	fmt.Println(bytes.NewBuffer(convertido))

	//EXEMPLO 2

	cachorro2 := map[string]string{
		"nome":  "Jack",
		"raca":  "Labrador",
		"idade": "5",
	}

	convertido2, error := json.Marshal(cachorro2)

	if error != nil {
		fmt.Println("Erro ao converter para JSON:", error)
	}

	// o resultado é um slice de bytes, então precisamos converter para string ou usar bytes.NewBuffer para imprimir
	fmt.Println(bytes.NewBuffer(convertido2))
}
