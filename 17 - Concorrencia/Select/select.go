package main

import (
	"fmt"
	"time"
)

func main() {

	canal := make(chan string)
	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			canal <- "Olá Mundo! Canal 1"
		}

	}()

	canal2 := make(chan string)
	go func() {
		for {
			time.Sleep(time.Second * 2)
			canal2 <- "Olá Mundo! Canal 2"
		}
	}()

	for {

		// O select é utilizado para esperar por múltiplas operações de canal. Ele bloqueia até que uma das operações de canal esteja pronta para ser executada. Quando uma operação de canal está pronta, o select executa a operação correspondente e continua a execução do programa.
		select {
		case msg1 := <-canal:
			fmt.Println(msg1)
		case msg2 := <-canal2:
			fmt.Println(msg2)
		}
	}
}
