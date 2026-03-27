package main

import "fmt"

func main() {
	numero := 10

	// Estrutura de controle if-else basica
	if numero > 15 {
		fmt.Println("O número é maior que 15")
	} else {
		fmt.Println("O número é menor ou igual a 15")
	}

	// Estrutura de controle if-else com declaração
	if numero2 := 5; numero2 > 0 {
		fmt.Println("é maior que zero")
	} else if numero2 == 0 {
		fmt.Println("é igual a zero")
	} else {
		fmt.Println("é menor que zero")
	}

}
