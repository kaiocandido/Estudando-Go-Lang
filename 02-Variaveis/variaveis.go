package main

import (
	"fmt"
)

func main() {
	var variavel1 string = "variável 1"
	variavel2 := "variável 2"

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	//declarando mais de uma variável utilizando o var
	var (
		varivel3 string = "variável 3"
		varivel4 int    = 22
	)

	fmt.Println(varivel3)
	fmt.Println(varivel4)

	//declarando mais de uma viariável sem o var, utilizando o operador de atribuição curta
	variavel5, variavel6 := "variável 5", 22
	fmt.Println(variavel5)
	fmt.Println(variavel6)

	// criando costantes

	const constante1 string = "constante 1"
	constante2 := "constante 2"

	fmt.Println(constante1)
	fmt.Println(constante2)

}
