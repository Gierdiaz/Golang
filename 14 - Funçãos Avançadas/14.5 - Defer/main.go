package main

import ("fmt")

func function1() {
	fmt.Println("Executando a função 1")
}

func function2() {
	fmt.Println("Executando a função 2")
}

func alunoAprovado(n1, n2 float32) bool {
	defer fmt.Println("Média calculada. Resultado será retornado")
	fmt.Println("Entrando na função para verificar se o aluno está aprovado")

	media := (n1 + n2) / 2

	if media >= 6 {
		
		return true
	}

	return false
}

func main() {
	defer function1()
	function2()

	fmt.Println(alunoAprovado(7,8))
}	