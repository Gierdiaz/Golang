package servidor

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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
func BuscarUsuarios (w http.ResponseWriter, r *http.Response) {
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
func BuscarUsuario (w http.ResponseWriter, r *http.Response) {

}