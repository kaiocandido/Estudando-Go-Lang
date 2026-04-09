package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type User struct {
	Name  string
	Email string
}

func main() {

	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		u := User{Name: "João", Email: "joao@example.com"}
		templates.ExecuteTemplate(w, "home.html", u)
	})

	fmt.Println("Executando na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", nil))

}
