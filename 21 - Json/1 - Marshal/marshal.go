package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type Animal struct {
	Nome   string `json:"nome"` 
	Idade  string `json:"idade"`
	Raca   string `json:"raça"`
}

func main() {

	cachorro := Animal{"Lucky", "2", "Dálmatas"}
	//fmt.Println(cachorro)

	cachorroJson, erro := json.Marshal(cachorro)

	if erro != nil {
		log.Fatal(erro)
	}

	//Duas formas de demonstrar o resultado json.

	//irá retornar um slice de bytes. Use 'string()'
	fmt.Println(string(cachorroJson))

	fmt.Println(bytes.NewBuffer(cachorroJson))

	//-----------------------------------------

	gato := map[string]string {"nome": "Argus", "raça": "vira-lata"}

	gatoJson, erro := json.Marshal(gato)

	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(bytes.NewBuffer(gatoJson))


}