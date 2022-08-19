package main

import ("fmt")

//Duas forma de usar o switch.

func diaSemana(numero int) string {
	switch numero {
		case 1:
			return "Domingo"
		case 2:
			return "Segunda"
		case 3:
			return "Terça"
		case 4:
			return "Quarta"
		case 5:
			return "Quinta"
		case 6: 
			return "Sexta"
		case 7:
			return "Sábado"
		default:
			return "Número inválido"	
	}

}

func mes(num int) string {
	switch {
		case num == 1:
			return "Janeiro"
		case num == 2:
			return "Fevereiro"
		case num == 3:
			return "Março"
		case num == 4:
			return "Abril"
		case num == 5:
			return "Maio"
		case num == 6: 
			return "Junho"
		case num == 7:
			return "Julho"
		case num == 8:
			return "Agosto"
		case num == 9:
			return "Setembro"
		case num == 10:
			return "Outubro"
		case num == 11:
			return "Novembro"
		case num == 12:
			return "Dezembro"
		default:
			return "Número inválido"	
	}

}

func main() {
	day := diaSemana(1)
	fmt.Println(day)

	month := mes(7)
	fmt.Println(month)
	
}