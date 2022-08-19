package main

import ("fmt")

func main() {
	//map: key value sempre do mesmo tipo
	usuario := map[string]string {
		"nome": "pedro",
		"sobrenome": "silva",
	}

	fmt.Println(usuario["nome"])

	user := map[string]map[string]string {
		"Esutdantes": {
			"Primeiro Aluno": "Arugs",
			"Segunda Aluno": "José",
		},

		"Curso": {
			"Tecnólogo": "Análise e Desenvolvimento de Sistema",
			"Bacharelado": "Segurança da Informação",
			"Licenciatura": "Informática",
		},
	}

	fmt.Println(user)
	delete(user,"Curso")
	fmt.Println(user)
	

	user["Felino"] = map[string]string {
		"A": "Gatinhos",
		"B": "Gatinhos",
	}

	fmt.Println(user)
	
}