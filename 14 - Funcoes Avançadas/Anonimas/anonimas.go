package main

import "fmt"

func main() {
	// Função anônima auto-invocada
	func() {
		fmt.Println("Olá, sou uma função anônima auto-invocada!")
	}()
}
