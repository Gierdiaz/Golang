package main

import ("fmt")

func main() {
	func (texto string) {
		fmt.Println(texto)
	}("Olá Mundo")

	// -------------------
	
	func () {
		fmt.Println("Olá Mundo")
	}()

	// -------------------

	retorno := func (txt string) string {
		return fmt.Sprintf("Recebido %s", txt)
	}("Passando parâmetro")

	fmt.Println(retorno)


}