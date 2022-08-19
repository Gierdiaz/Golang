package main

import ("fmt")

type usuario struct {
	nome string
	sobrenome string
	
}

type endereco struct {
	logradouro string
	numero int
	usuario
}

func main() {
	u := usuario {
		nome: "Allison",
		sobrenome: "Luis",
	}
	fmt.Println(u.sobrenome)

	e := endereco {
		usuario: usuario {
			nome: "José",
			sobrenome: "gatinho",
		},

		logradouro: "não sei",
		numero: 1,
		
	} 
	fmt.Println(e)


}