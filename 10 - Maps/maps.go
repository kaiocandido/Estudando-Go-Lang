package main

import "fmt"

func main() {
	fmt.Println("Maps")

	// Criando um mapa para armazenar informações de um usuário
	usuario := map[string]string{
		"nome":      "João",
		"sobrenome": "Silva",
		"email":     "kaio.candf@gmail.com",
	}

	fmt.Println(usuario)         // Imprimindo o mapa completo
	fmt.Println(usuario["nome"]) // Acessando o valor da chave "nome"

	usuario2 := map[string]map[string]string{
		"usuario1": {
			"nome":      "Maria",
			"sobrenome": "Santos",
		},
	}
	fmt.Println(usuario2)   // Imprimindo o mapa aninhado
	delete(usuario, "nome") // Removendo a chave "nome" do mapa usuario

	fmt.Println(usuario) // Imprimindo o mapa após a remoção

	usuario2["signo"] = map[string]string{
		"signo1": "Áries",
	}
	fmt.Println(usuario2) // Imprimindo o mapa após adicionar um novo mapa aninhado

}
