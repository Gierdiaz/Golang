package main

import (
		"log"
		"net/http"
)

func main() {

	//Protocolo de comunicação - Base de comunicação web

	//Cliente - Servidor

	//Request - Response

	//Rotas
	// URI - Identificador de Recurso
	// Método - GET, POST, PUT, DELETE

	http.HandleFunc("/home", home)

	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Campo de usuários"))
	})

	log.Fatal(http.ListenAndServe(":5000", nil))
}
// duas formas de presentar a função

func home (w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Campo principal"))
}