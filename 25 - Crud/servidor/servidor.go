package servidor

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type usuario struct {
	ID    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, err := ioutil.ReadAll(r.Body)
	
	if err != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição"))
		return
	}

	var usuario usuario

	if err = json.Unmarshal(corpoRequisicao, &usuario); err != nil {
		w.Write([]byte("Erro ao converter o usuário para struct"))
		return
	}

	fmt.Println(usuario)

	db, err := banco.Conectar()

	if err != nil {
		w.Write([]byte("Erro ao conectar no banco de dados"))
		return
	}
	
	defer db.Close()

	// INSERT INTO USUARIOS (nome, email) values ("nome", "email")

	//PREPARE STATEMENT

	statement, err := db.Prepare("insert into usuarios (nome, email) values (?, ?)")

	if err != nil {
		w.Write([]byte("Erro ao criar o statement"))
		return
	}

	defer statement.Close()

	insercao, err := statement.Exec(usuario.Nome, usuario.Email)

	if err != nil {
		w.Write([]byte("Erro ao executar o statement"))
		return
	}

	idInserido, err := insercao.LastInsertId()

	if err != nil {
		w.Write([]byte("Erro ao obter o id inserido"))
		return
	}

	// STATUS CODES
	w.WriteHeader(http.StatusCreated)

	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso! Id: %d", idInserido )))
	
}

// BuscarUsuarios traz todos os usuários salvos no banco
func BuscarUsuarios (w http.ResponseWriter, r *http.Request) {

	db, err := banco.Conectar() 

	if err != nil {
		w.Write([]byte("Erro ao conectar o banco de dados"))
	}

	defer db.Close()

	// SELECT * FROM USUARIOS

	linhas, err := db.Query("select * from usuarios")

	if err != nil {
		w.Write([]byte("Erro ao buscar usuários"))
	}

	defer linhas.Close()

	var usuarios []usuario

	for linhas.Next() {
		var usuario usuario

		if err := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); err != nil {
			w.Write([]byte("Erro ao escanear o usuário"))

			return
		}

		usuarios = append(usuarios, usuario)
	}

	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(usuarios); err != nil {
		w.Write([]byte("Erro ao converte os usuários para json"))

		return
	}
}

// BuscarUsuário traz um usuário específico salvo no banco de dados
func BuscarUsuario (w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, err := strconv.ParseUint(parametros["id"], 10, 64)

	if err != nil {
		w.Write([]byte("Erro ao converter o parâmetro para inteiro"))

		return
	}

	db, err := banco.Conectar()

	if err != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados"))

		return
	}

	linha, err := db.Query("Select * from usuarios where id = ?", ID)

	if err != nil {
		w.Write([]byte("Erro oa buscar o usuário"))

		return
	}

	var usuario usuario
	if linha.Next() {
		if err := linha.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); err != nil {
			w.Write([]byte("Erro ao escanear o usuario"))

			return
		}
	}

	if err := json.NewEncoder(w).Encode(usuario); err != nil {
		w.Write([]byte("Erro ao converter o usuário para json"))

		return
	}

}

// AtualizarUsuario altera os dados de um usuário no banco de dados
func AtualizarUsuario (w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, err := strconv.ParseUint(parametros["id"], 10, 64)

	if err != nil {
		w.Write([]byte("Erro ao converter o parâmetro para inteiro"))

		return
	}

	corpoRequisicao, err := ioutil.ReadAll(r.Body) 

	if err != nil {
		w.Write([]byte("Erro ao ler o corpo da requisição"))

		return
	}

	var usuario usuario 

	if err := json.Unmarshal(corpoRequisicao, &usuario); err != nil {
		w.Write([]byte("Erro ao converter o usuário para struct"))

		return
	}

	db, err := banco.Conectar()

	if err != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados"))

		return
	}

	defer db.Close()

	statement, err := db.Prepare("update usuarios set nome = ?, email = ? where id = ?")

	if err != nil {
		w.Write([]byte("Erro ao criar o statement"))
	}

	defer statement.Close()

	if _, err := statement.Exec(usuario.Nome, usuario.Email, ID); err != nil {
		w.Write([]byte("Erro ao atualizar o usuarário"))

		return
	}

	w.WriteHeader(http.StatusNoContent)


}

// DeletaUsuario remove um usuário do banco de dados
func DeletaUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, err := strconv.ParseUint(parametros["id"], 10, 64)

	if err != nil {
		w.Write([]byte("Erro ao converter o parâmetro para interio"))

		return
	}

	db, err := banco.Conectar() 

	if err != nil {
		w.Write([]byte("Erro ao conectar ao banco de dados"))

		return
	}

	defer db.Close()

	statement, err := db.Prepare("delete from usuarios where id = ?")

	if err != nil {
		w.Write([]byte("Erro ao criar o statement"))

		return
	}

	defer statement.Close()

	if _, err := statement.Exec(ID); err != nil {
		w.Write([]byte("Erro ao deletar o usuário"))

		return
	}

	w.WriteHeader(http.StatusNoContent)
}