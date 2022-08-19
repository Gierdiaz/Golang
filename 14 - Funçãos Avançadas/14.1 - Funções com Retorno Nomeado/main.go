package main

import ("fmt")

// (soma int, subtr int) retorno da função
func calculoMatematico(n1 int,n2 int) (soma int, subtr int) {
	soma = n1 + n2
	subtr = n1 - n2
	return
} 


func main() {
	
	soma, subrt := calculoMatematico(7,3)
	fmt.Println(soma, subrt)
	
}