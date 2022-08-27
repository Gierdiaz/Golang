package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type Usuario struct {
	Nome  string
	Email string
}

func main () {

	fmt.Println("Escutando na porta 5000")

	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/home", home)

	log.Fatal(http.ListenAndServe(":5000", nil))

}


func home (w http.ResponseWriter, r *http.Request) {
	
	user := Usuario {Nome: "Argus", Email: "argus@gatinho.com.br"}

	templates.ExecuteTemplate(w, "home.html", user)

}	