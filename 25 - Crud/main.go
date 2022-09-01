package main

import (
	"fmt"
	"log"
	"net/http"
	"crud/servidor"

	"github.com/gorilla/mux"
)

func main() {
	//CRUD = CREATE, READ, UPDATE, DELETE

	//CREATE - POST
	//READ - GET
	//UPDATE - PUT
	//DELETE - DELETE

	router := mux.NewRouter()
	router.HandleFunc("/usuarios", servidor.CriarUsuario).Methods(http.MethodPost)



	fmt.Println("Escutando na porta 4000")
	log.Fatal(http.ListenAndServe(":4000", router))
}