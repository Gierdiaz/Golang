package main

import ("fmt")

func inverterNumero(numero int) int {
	return numero * - 1
}

func inverterSinalPonteiro(numero *int) {
	*numero = *numero * - 1
}

func main() {
	numero := 20
	numeroInvertido := inverterNumero(numero)
	fmt.Println(numeroInvertido)

	novoNumero := 40
	inverterSinalPonteiro((&novoNumero))
	fmt.Println(novoNumero)
}