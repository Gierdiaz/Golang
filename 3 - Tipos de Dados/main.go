package main

import (
	"errors"
	"fmt"
)

func main() {
	//int, int16, int32, int64, uint(positivos) //bits

	var numero1 int16 = 100
	fmt.Println(numero1)

	var numero2 uint32 = 100
	fmt.Println(numero2)

	var numero3 float32 = 1.5
	fmt.Println(numero3)

	var numero4 float64 = 15646815.6546
	fmt.Println(numero4)

	var str1 string = "A"
	fmt.Println(str1)

	//inferência de tipo
	str2 := "B"
	fmt.Println(str2)

	//valor zero
	var texto string 
	fmt.Println(texto)

	var booleano1 bool 
	fmt.Println(booleano1)

	//pacote. Error muito usado em go
	var erro error = errors.New("Aceito")
	fmt.Println(erro)
}