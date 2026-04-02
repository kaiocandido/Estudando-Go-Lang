package main

import "fmt"

func clousure() func() {
	texto := "Olá, mundo!"

	funcao := func() {
		fmt.Println(texto)
	}

	return funcao
}

func main() {
	texto := "ainda não acessei o texto do main"
	fmt.Println(texto)

	funcnova := clousure()
	funcnova()
}
