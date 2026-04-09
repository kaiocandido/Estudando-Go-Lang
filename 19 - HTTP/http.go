package main

import (
	"log"
	"net/http"
)

// http é um protocolo de comunicação entre clientes e servidores.
// Ele é utilizado para transferir dados na web, como páginas HTML, imagens, vídeos, etc. Em Go, o pacote "net/http"
// fornece funcionalidades para criar servidores HTTP e fazer requisições HTTP.

// Rotas são os caminhos que os clientes podem acessar em um servidor HTTP. Por exemplo, se um cliente acessar
// "http://localhost:8080/hello", a rota seria "/hello". O servidor pode responder a essa rota com uma mensagem ou
// um recurso específico.

// URI (Uniform Resource Identifier) é um identificador único para um recurso na web. Ele pode ser uma URL
// (Uniform Resource Locator) e Metodo HTTP é a ação que o cliente deseja realizar em um recurso.
// Os métodos HTTP mais comuns são GET, POST, PUT, DELETE, etc.

func main() {
	// Neste exemplo, estamos registrando a rota "/home" e a função anônima que escreve "Hello world" na resposta.
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	})

	http.HandleFunc("/usuarios", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Paginas de Usuarios"))
	})

	// http.HandleFunc é uma função que registra uma rota e a função que será executada quando essa rota for acessada.
	log.Fatal(http.ListenAndServe(":5000", nil))

}
