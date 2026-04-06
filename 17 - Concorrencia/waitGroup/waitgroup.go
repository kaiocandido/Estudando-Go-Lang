package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var waitGroupp sync.WaitGroup
	waitGroupp.Add(2) // Adiciona 2 goroutines ao WaitGroup

	go func() {
		hello("Hello, World!")
		waitGroupp.Done()
	}()

	go func() {
		hello("Programação concorrente em Go!")
		waitGroupp.Done()
	}()

	waitGroupp.Wait() // Aguarda as goroutines terminarem

}

func hello(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(1 * time.Second)
	}

}
