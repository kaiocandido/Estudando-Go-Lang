package servidor

import (
	banco "crud/Banco"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type usuarioStruct struct {
	ID    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

// Inserir um novo usuário no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição!"))
		return
	}

	var usuario usuarioStruct

	if err = json.Unmarshal(corpoRequisicao, &usuario); err != nil {
		w.Write([]byte("Falha ao converter o corpo da requisição para o formato JSON!"))
		return
	}

	db, err := banco.Conectar()

	if err != nil {
		w.Write([]byte("Falha ao conectar com o banco de dados!"))
		return
	}

	defer db.Close()

	statenment, err := db.Prepare("insert into usuarios (nome, email) values (?, ?)")

	if err != nil {
		w.Write([]byte("Falha ao criar o statenment!"))
		return
	}

	defer statenment.Close()

	insercao, err := statenment.Exec(usuario.Nome, usuario.Email)

	if err != nil {
		w.Write([]byte("Falha ao executar o statenment!"))
		return
	}

	Idinserido, err := insercao.LastInsertId()

	if err != nil {
		w.Write([]byte("Falha ao obter o ID do usuário inserido!"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Usuário inserido com sucesso! ID: " + string(Idinserido)))

}

// Buscar todos os usuários do banco de dados
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	db, err := banco.Conectar()

	if err != nil {
		w.Write([]byte("Erro ao conectar com banco"))
	}

	defer db.Close()

	linhas, err := db.Query("select * from usuarios")

	if err != nil {
		w.Write([]byte("Erro ao buscar os usuários!"))
	}

	defer linhas.Close()

	var usuarios []usuarioStruct

	for linhas.Next() {
		var usuario usuarioStruct

		// metodo Scan é utilizado para ler os dados de cada linha retornada pela consulta e armazená-los na struct usuarioStruct

		if err = linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); err != nil {
			w.Write([]byte("Erro ao scanear o usuario"))
			return
		}

		usuarios = append(usuarios, usuario)
	}

	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(usuarios); err != nil {
		w.Write([]byte("Erro ao converter usuarios para json"))
		return
	}
}

// Buscar um usuário específico do banco de dados pelo ID
func BuscarUsuariosPorId(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, err := strconv.ParseUint(parametros["id"], 10, 32)

	if err != nil {
		w.Write([]byte("Erro ao converter para inteiro!"))
		return
	}

	db, err := banco.Conectar()

	if err != nil {
		w.Write([]byte("Erro ao conectar ao banco!"))
		return
	}

	linha, err := db.Query("select * from usuarios where id = ?", ID)

	if err != nil {
		w.Write([]byte("Erro ao buscar usuario!"))
	}

	var usuario usuarioStruct

	if linha.Next() {
		if err := linha.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); err != nil {
			w.Write([]byte("Erro ao escanear usuario!"))
			return
		}
	}

	if err := json.NewEncoder(w).Encode(usuario); err != nil {
		w.Write([]byte("Erro ao fazer o encode do usuario!"))
		return
	}

}

// Atualizar usuario no banco de dados
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	paratros := mux.Vars(r)

	ID, err := strconv.ParseUint(paratros["id"], 10, 32)

	if err != nil {
		w.Write([]byte("Erro ao converter parametro para inteiro"))
		return
	}

	corpoDaReq, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.Write([]byte("Erro ao ler o corpo da mensagem"))
		return
	}

	var usuario usuarioStruct

	if err := json.Unmarshal(corpoDaReq, &usuario); err != nil {
		w.Write([]byte("Erro para converter para struct"))
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		w.Write([]byte("Erro ao conectar no banco de dados!"))
		return
	}

	defer db.Close()

	statement, err := db.Prepare("update usuarios set nome = ?, email = ? where id = ?")
	if err != nil {
		w.Write([]byte("Erro ao criar o statement"))
		return
	}

	defer statement.Close()

	if _, err := statement.Exec(usuario.Nome, usuario.Nome, ID); err != nil {
		w.Write([]byte("Erro ao atualizar o ID"))
		return
	}

	w.WriteHeader(http.StatusNoContent)

}

// Deletar Usuario especifico pelo ID
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {

	parametros := mux.Vars(r)

	ID, err := strconv.ParseUint(parametros["id"], 10, 32)

	if err != nil {
		w.Write([]byte("Erro converter para inteiro"))
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		w.Write([]byte("Erro ao se conectar com banco dedados"))
		return
	}

	defer db.Close()

	statement, err := db.Prepare("delete from usuarios where id = ?")
	if err != nil {
		w.Write([]byte("Erro ao deletar usuario!"))
		return
	}

	defer statement.Close()

	if _, err := statement.Exec(ID); err != nil {
		w.Write([]byte("Erro ao deletar usuario!!"))
		return
	}

	w.WriteHeader(http.StatusNoContent)

}
