package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string) 
	go Write("Hello, World", channel)

	fmt.Println("Depois da função escrever ser executada")
	for {
		message, open := <- channel // recebendo um valor 
		if !open {
			break
		}
		fmt.Println(message)
	}
	fmt.Println("end of program")
		
}

func Write(txt string, channel chan string) {
	time.Sleep(time.Second * 3)
	for i := 0; i < 5; i++ {
		channel <- txt //mandando um valor para dentro do canal
		time.Sleep(time.Second)
	}

	close(channel)
}