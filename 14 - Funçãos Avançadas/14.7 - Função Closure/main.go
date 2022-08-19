package main

import ("fmt")

func closure() func() {
	texto := "Dentro da função closure"

	function := func() {
		fmt.Println(texto)
	}

	return function
}

func main () {
	texto := "Dentro da função main"
	fmt.Println(texto)

	newFunction := closure()
	newFunction()
}