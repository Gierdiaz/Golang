package main

import (
	"fmt"
	"log"
	"os"
	"linha-de-comando/app"
)

func main() {
	fmt.Println("Ponto de partida")

	application := app.Gerar()
	erro := application.Run(os.Args)
	if erro != nil {
		log.Fatal(erro)
	}
}