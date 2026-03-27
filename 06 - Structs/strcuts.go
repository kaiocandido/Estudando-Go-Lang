package main

import (
	"fmt"
)

type usuario struct {
	nome  string
	idade int
}

func main() {
	fmt.Println("Hello, World!")

	//maneira de declarar uma struct
	var u1 usuario
	u1.nome = "João"
	u1.idade = 30

	//outra maneira de declarar uma struct
	u2 := usuario{
		nome:  "Maria",
		idade: 25,
	}
	fmt.Println(u1.idade)
	fmt.Println(u2.idade)
}
