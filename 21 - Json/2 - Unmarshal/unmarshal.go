package main

import (
	"fmt"
	"log"
	"encoding/json"
	_"bytes"
)

type Animal struct {
	Nome  string `json:"nome"`
	Idade string `json:"idade"`
	Raca  string `json:"raça"`
}


func main() {
	cachorroJSON := `{"nome": "Lucky", "idade": "3", "raça": "Dálmatas"}`

	var cachorro Animal

	if erro := json.Unmarshal([]byte(cachorroJSON), &cachorro); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cachorro)

	// =========================================================================

	gatoJSON := `{"nome": "José", "idade": "2", "raça": "siamês"}`

	gato := make(map[string]string)

	if erro := json.Unmarshal([]byte(gatoJSON), &gato); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(gato)
}