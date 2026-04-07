package main

import (
	"fmt"
	"time"
)

func main() {
	go hello("Hello, World!") // goroutine principal é um metodo que é executado quando o programa é iniciado, e é responsável por iniciar outras goroutines.
	hello("Programação concorrente em Go!")

}

func hello(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(1 * time.Second)
	}

}
