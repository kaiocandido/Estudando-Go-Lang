package main

import (
	"fmt"
)

func main() {
	// Operadores Aritméticos
	// + - * / %

	soma := 10 + 5
	subtracao := 10 - 5
	multiplicacao := 10 * 5
	divisao := 10 / 5
	modulo := 10 % 3

	fmt.Println(soma)
	fmt.Println(subtracao)
	fmt.Println(multiplicacao)
	fmt.Println(divisao)
	fmt.Println(modulo)

	// Operadores de Atribuição
	// = :=
	var x int = 10
	y := 2
	fmt.Println(x)
	fmt.Println(y)

	// Operadores de Relacionais
	// == != > < >= <=
	fmt.Println(10 == 5)
	fmt.Println(10 != 5)
	fmt.Println(10 > 5)
	fmt.Println(10 < 5)
	fmt.Println(10 >= 5)

	// Operadores Lógicos
	// && || !
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	// Operadores Unários
	// ++ --
	z := 10
	z++
	fmt.Println(z)
	z--
	fmt.Println(z)
}
