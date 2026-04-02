package main

import "fmt"

func one() {
	fmt.Println("Executando a função one")
}

func two() {
	fmt.Println("Executando a função two")
}

func main() {
	defer one() //adiar a execução da função one até o final da função main
	two()
}
