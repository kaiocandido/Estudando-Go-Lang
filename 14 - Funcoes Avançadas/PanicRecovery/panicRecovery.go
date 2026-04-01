package main

import "fmt"

func recuperar() {
	if r := recover(); r != nil {
		fmt.Println("Recuperando do erro:", r)
	}
}

func notas(n1, n2 float64) bool {
	defer recuperar() //defer é usado para garantir que a função de recuperação seja chamada mesmo que ocorra um panic
	media := (n1 + n2) / 2

	if media > 6 {
		return true

	} else if media < 6 {
		return false
	}

	panic("A média é  6") //gera um erro e interrompe a execução do programa

}

func main() {
	fmt.Println("Iniciando o programa")
	fmt.Println(notas(6, 6))
}
