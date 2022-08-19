package main

import ("fmt")

func soma(num ... int) int {
	total := 0
	for _, v := range num {
		total = total + v	
	}
	return total
}

//Imprimindo string + int
func Escrever(texto string, numero ... int) {
	for _, v := range numero {
		fmt.Println(texto, v)
	}
}

func main() {
	totalSoma := soma(1,2,3,4,5,6,7,8,9)
	fmt.Println(totalSoma)

	Escrever("Olá Mundo!", 1,2,3)
}