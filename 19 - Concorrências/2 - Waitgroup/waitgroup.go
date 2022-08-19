package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(3) //Quantidade de goroutines que estarão rodando ao mesmo tempo.

	go func () {
		Write("Hello World")
		waitGroup.Done() //-1
	}()

	go func () {
		Write("Programando em go")
		waitGroup.Done() //-1
	}()

	waitGroup.Wait() 

}

func Write(txt string) {
	fmt.Println(txt)

	for i := 0; i < 5; i++ {
		fmt.Println(txt)
		time.Sleep(time.Second)
	}
}