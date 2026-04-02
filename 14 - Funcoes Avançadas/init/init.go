package main

import "fmt"

var n int

func init() {
	n = 10
	// inicia primeiro, antes do main
	// usada para inicializar variáveis, abrir conexões, etc.
	fmt.Println("Init function called")

}
func main() {
	fmt.Println("Hello, World!")
	fmt.Println("Value of n:", n)
}
