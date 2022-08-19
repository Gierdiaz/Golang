package main

import ("fmt")

func main() {
	channel := make(chan string, 1) //capacidade de um

	channel <- "Hello World"

	message := <- channel 
	fmt.Println(message)
}