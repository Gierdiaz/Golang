/*Função (variável, struct) nomeFunção() {} */
package main

import ("fmt")

//struct
type usuario struct {
	idade int
	nome string
}

//método
func (user usuario) salvar() {
	fmt.Printf("Salvando os dados do usuário %s no banco de dados\n", user.nome)
}

//método. Verificando se é maior de idade
func (user usuario) maiorIdade() bool {
	return user.idade >= 18
}

func (user *usuario) fazerAniversario() {
	user.idade = user.idade + 1
	
}

func main() {
	rafael := usuario {
		idade: 19,
		nome: "Rafael",
	}

	mateus := usuario {
		idade: 16,
		nome: "Mateus",
	}

	miguel := usuario {
		idade: 17,
		nome: "Miguel",
	}

	//chamando o struct
	fmt.Println(rafael)

	//chamando o método
	rafael.salvar()
	mateus.salvar()
	miguel.salvar()

	//chamando e verificando
	maiorIdade := rafael.maiorIdade()
	fmt.Println(maiorIdade)
	
	//alteração no valor
	mateus.fazerAniversario()
	fmt.Println(mateus.idade)
}