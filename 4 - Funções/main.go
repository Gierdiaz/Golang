package main

import "fmt"

func somar(n1 int, n2 int) int {
	return n1 + n2
}

func calculosMatematicos(t1 int, t2 int) (int, int, int, int) {
	soma := t1 + t2
	subtracao := t1 - t2
	multiplicacao := t1 * t2
	divisao := t1/t2
	return soma, subtracao, multiplicacao, divisao 

}

func main() {
	soma := somar(2, 2)
	fmt.Println(soma)

	//resulSoma. Caso você não queira usar a variável resulSoma
	_, resulSub, resulMult, resulDiv := calculosMatematicos(10,5)
	fmt.Println(resulSub, resulMult, resulDiv)

	//nível 1
	var j = func() {
		fmt.Println("Função j")
	}
	j()

	//nível 2
	var f = func(txt string) {
		fmt.Println(txt)
	}
	f("função f")

	//nível 3
	var k = func(str string) string {
		return str
	}
	resultado := k("função k")
	fmt.Println(resultado)

}