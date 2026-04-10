package main

import (
	servidor "crud/Servidor"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/usuarios", servidor.CriarUsuario).Methods("POST")

	fmt.Println("Servidor rodando na porta 5000")

	log.Fatal(http.ListenAndServe(":5000", router))
}
