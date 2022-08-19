package main

import ("fmt")

type pessoa struct {
	nome string
	idade int
	identidade string
	cpf string

}

type estudante struct {
	matricula int
	faculdade string
	curso string
	pessoa
	
}

func main() {
	p := pessoa{
		nome:       "João",
		idade:      29,
		identidade: "30.590.780-9",
		cpf:        "162.356.789-30",
	}

	e := estudante{
		matricula: 202104732998,
		faculdade: "Ecole Polytecnique",
		curso:     "Engenharia Aeroespacial",
		

		pessoa: pessoa{
			nome:       "João",
			idade:      29,
			identidade: "30.590.780-9",
			cpf:        "162.356.789-30",
		},
	}
	fmt.Println(e.faculdade)
	fmt.Println(p)
}