package main

import ("fmt")

func main() {
	//aritméticos
	soma := 1 + 1
	sub := 1 - 1
	mult := 1 * 1
	div := 1 / 1

	fmt.Println(soma, sub, mult, div)

	//atribuição
	var variavel1 string = "string"
	fmt.Println(variavel1)

	//operadores relacionais
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2 )
	fmt.Println(1 != 2)
	
	//Operadores lógicos
	fmt.Println("-------------")
	verdade, falso := true, false
	fmt.Println(verdade && falso)
	fmt.Println(verdade || falso)
	fmt.Println(!verdade)
	fmt.Println(!falso)

	//operadores unitários
	numero := 10
	numero = numero + 1
	fmt.Println(numero)


}