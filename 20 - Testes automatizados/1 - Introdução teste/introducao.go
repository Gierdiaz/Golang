package main

import (
		"introducao-testes/enderecos"
		"fmt"

)


func main() {
	TipoEndereco := enderecos.TipoEndereco("rua Olinto")
	fmt.Println(TipoEndereco)
}