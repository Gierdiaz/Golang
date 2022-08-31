package main

import (
	"database/sql"
	"fmt"
	"log"
	_"github.com/go-sql-driver/mysql"
)

func main() {
	stringConexao := "usuario:senha@/banco"

	db, err := sql.Open("mysql", stringConexao)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conexão está aberta.")

	linhas, err := db.Query("select * from usuarios")

	if err != nil {
		log.Fatal(err)
	}

	defer linhas.Close()

	fmt.Println(linhas)
}