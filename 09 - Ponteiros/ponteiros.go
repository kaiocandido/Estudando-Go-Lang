package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var a int = 10
	var b int = a

	fmt.Println("Valor de a:", a)
	fmt.Println("Valor de b:", b)
	a++
	fmt.Println("Valor de a após incremento:", a)
	fmt.Println("Valor de b (sem alteração):", b)

	// PONTEIRO É UMA REFERENCIA DE MEMORIA

	var c int = 100
	var ponteiro *int
	fmt.Println(c, ponteiro) // Imprime o valor de c e o endereço de memória armazenado em ponteiro

	c = 200
	ponteiro = &c             // Atualiza o ponteiro para o novo endereço de c
	fmt.Println(c, *ponteiro) // Imprime o novo valor de c e o novo endereço de memória armazenado em ponteiro

}
