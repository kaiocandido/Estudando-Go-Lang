package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	urlConexão := "golang:golang@/deevbook?charset=utf8&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", urlConexão)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Coenxão estabelecida com sucesso!")

	linhas, err := db.Query("select * from usuarios")

	if err != nil {
		log.Fatal(err)
	}

	defer linhas.Close()

}
