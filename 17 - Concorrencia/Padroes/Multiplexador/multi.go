package main

import (
	"fmt"
	"time"
)

func main() {
	canal := multiplexar(escrever("Ola"), escrever("Mundo!"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

func multiplexar(canalDeEntrada1, canalDeEntrada2 <-chan string) <-chan string {
	canalDeSaida := make(chan string)

	go func() {
		for {
			select {
			case msg := <-canalDeEntrada1:
				canalDeSaida <- msg
			case msg := <-canalDeEntrada2:
				canalDeSaida <- msg
			}
		}
	}()
	return canalDeSaida
}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf(texto)
			time.Sleep(time.Millisecond * 500)
		}
	}()
	return canal
}
