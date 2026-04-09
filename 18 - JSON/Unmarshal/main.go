package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Idade int    `json:"idade"`
	Raca  string `json:"raca"`
}

func main() {

	cachorroEmJSON := `{"nome": "Rex", "idade": 5, "raca": "Labrador"}`

	var c cachorro

	// Unmarshal é a função que converte um JSON em uma struct Go.
	if erro := json.Unmarshal([]byte(cachorroEmJSON), &c); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(c)

	// Também é possível converter um JSON em um map[string]string, por exemplo.

	cachorro2EmJSON := `{"nome": "Tobi", "raca": "Labrador"}`

	c2 := make(map[string]string)

	if erro := json.Unmarshal([]byte(cachorro2EmJSON), &c2); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(c2)
}
