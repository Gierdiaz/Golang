package main

import ("fmt")

func main() {
	var variavel1 int = 10
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)

	variavel1 = variavel1 + 1
	fmt.Println(variavel1, variavel2)

	//ponteiro é uma referência de memória (100 0xc0000ac080)
	//jogando o endereço de memória onde o valor 100 está armazenada
	var variavel3 int = 100
	var ponteiro *int = &variavel3

	fmt.Println(variavel3, ponteiro)
	fmt.Println(variavel3, *ponteiro) //desreferênciação

}