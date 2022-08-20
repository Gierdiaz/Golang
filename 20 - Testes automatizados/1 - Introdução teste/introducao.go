package main

import (
		"fmt"
		"introducao-testes/enderecos"
)


func main() {
	TipoEndereco := enderecos.TipoEndereco("rua Olinto")
	fmt.Println(TipoEndereco)
}