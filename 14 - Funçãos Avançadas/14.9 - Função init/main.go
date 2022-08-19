package main

import ("fmt")

var n int

//ela executa primeiro do que a função main
func init() {
	fmt.Println("Executando a função init")

	n = 10
}

func main() {
	fmt.Println("Executando a função main")

	fmt.Println(n)
}