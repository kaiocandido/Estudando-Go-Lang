package main

import "fmt"

func main() {
	canal := make(chan string, 2) // Buffer de tamanho 2 significa que o canal pode armazenar até 2 mensagens sem bloquear a goroutine que envia as mensagens.
	canal <- "Olá Mundo"
	msg := <-canal
	fmt.Println(msg)

}
