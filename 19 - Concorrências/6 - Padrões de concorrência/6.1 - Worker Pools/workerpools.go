package main

import ("fmt")

func main() {
	tarefas := make(chan int, 45)
	resultado := make(chan int, 45)

	//acelerando o processo
	go Worker(tarefas, resultado)
	go Worker(tarefas, resultado)
	go Worker(tarefas, resultado)
	go Worker(tarefas, resultado)

	for i := 0; i < 45; i++ {
		tarefas <- i
	}
	close(tarefas)

	for i := 0; i < 45; i++ {
		resultado := <- resultado
		fmt.Println(resultado)
	}

}

func Worker(tarefas <-chan int, resultado chan<- int) {
	for numero := range tarefas {
		resultado <- fibonacci(numero)
	}
}

func fibonacci (p int) int {
	if p <= 1 {
		return p
	}

	return fibonacci(p - 2) + fibonacci(p - 1)
}