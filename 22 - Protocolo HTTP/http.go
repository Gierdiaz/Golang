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

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Estou no ar"))
	})

	log.Fatal(http.ListenAndServe(":5000", nil))
}