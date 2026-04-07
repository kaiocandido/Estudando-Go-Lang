package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)
	go hello("Olá mundo!", canal)

	for {
		// deadlock: o programa trava porque o canal não tem mais mensagens para ler, mas o loop continua tentando ler
		mensagem, aberto := <-canal

		// para evitar o deadlock, verificamos se o canal está aberto antes de tentar ler a mensagem
		if !aberto {
			break
		}

		fmt.Println(mensagem)
	}

	fmt.Println("Fim do programa!!")

}

func hello(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto
		time.Sleep(1 * time.Second)
	}

	close(canal) // fecha o canal para indicar que não haverá mais mensagens

}
