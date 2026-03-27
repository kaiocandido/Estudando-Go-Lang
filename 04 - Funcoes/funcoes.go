package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculos(n1 int, n2 int) (soma int, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	return soma, subtracao
}

func main() {
	resultado := somar(10, 20)
	println(resultado)

	var f = func() {
		fmt.Println("Função anônima")
	}

	f()

	cal, _ := calculos(10, 5)
	fmt.Println(cal)

}
