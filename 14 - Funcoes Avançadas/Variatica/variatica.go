package main

// Função variática: permite receber um número variável de argumentos
func somaNums(numeros ...int) int {
	soma := 0
	for _, num := range numeros {
		soma += num
	}
	return soma
}

func main() {
	totalSoma := somaNums(1, 2, 3, 4, 5)
	println("A soma dos números é:", totalSoma)
}
