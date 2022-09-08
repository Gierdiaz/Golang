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
	router.HandleFunc("/usuarios", servidor.BuscarUsuarios).Methods(http.MethodGet)
	router.HandleFunc("/usuarios/{id}", servidor.BuscarUsuario).Methods(http.MethodGet)
	router.HandleFunc("/usuarios/{id}", servidor.AtualizarUsuario).Methods(http.MethodPut)
	router.HandleFunc("/usuarios{id}", servidor.DeletaUsuario).Methods(http.MethodDelete)



	fmt.Println("Escutando na porta 4000")
	log.Fatal(http.ListenAndServe(":4000", router))
}