package banco

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // Importa o driver do MySQL para que ele seja registrado com o pacote "database/sql"
)

func Conectar() (*sql.DB, error) {
	stringDeConexao := "golang:golang@tcp(127.0.0.1:3306)/deevbook?charset=utf8&parseTime=True&loc=Local"

	db, err := sql.Open("mysql", stringDeConexao)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
