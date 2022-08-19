package main

import (
	"fmt"
	"time"
)

func main() {
	go Escrever("Hello World") //goroutine
	Escrever("Programando em go")
}

func Escrever(texto string) {
	fmt.Println(texto)

	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}