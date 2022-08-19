package main

import ("fmt")

func main() {
	//atribuição de variável e constante
	var variavel1 string = "varivavel 1"
	//Segue como as variáveis, porém substituindo por constantes
	const constante1 string = "constante 1"


	fmt.Println(constante1)
	fmt.Println(variavel1)

	variavel2 := "variavel 2"
	fmt.Println(variavel2)

	var (
		variavel3 string = "variavel 3"
		variavel4 int = 4
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "variavel5", "variavel6"
	fmt.Println(variavel5, variavel6)
}