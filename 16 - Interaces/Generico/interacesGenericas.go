package main

import "fmt"

// interface generica permite que uma função ou tipo aceite qualquer tipo de dado, sem restrições específicas. Isso é útil quando queremos criar funções ou estruturas que possam trabalhar com diferentes tipos de dados sem precisar definir um tipo específico para cada um.
func generica(a interface{}) {
	fmt.Println(a)
}

func main() {
	generica(42) // Passando um inteiro
	generica("Hello")
	generica(3.14)           // Passando um float
	generica([]int{1, 2, 3}) // Passando um slice de inteiros
	generica(true)           // Passando um booleano
	fmt.Println("Fim do programa")

}
