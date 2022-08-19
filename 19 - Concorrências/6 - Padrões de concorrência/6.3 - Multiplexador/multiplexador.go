package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main( ) {
	canal := Multiplexador(Escrever("Olá Mundo"), Escrever("Progamando em Go"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

func Multiplexador (canalEntrada1, canalEntrada2 <-chan string) <-chan string {
	canalSaida := make(chan string)

	go func () {
		for {
			select {
			case mensagem := <-canalEntrada1:
				canalSaida <- mensagem
			case mensagem := <- canalEntrada2:
				canalSaida <- mensagem

			}
		}
	}()

	return canalSaida
}

func Escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido %s", texto)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return canal
}