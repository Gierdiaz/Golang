package main

import (
	"fmt"
	"time"
)

func main() {
		i := 0

		for i < 3 {
			i++
			time.Sleep(time.Second)
			fmt.Println("Incrementando i", i)
			
		}

		// -------------------------------------

		for	j := 0; j <= 3; j++ {
			fmt.Println("Incrementando J", j)
			time.Sleep(time.Second)
		}

		// -------------------------------------

		nomes := [3]string {"José", "Argus", "Gudi"}

		for i, v := range nomes {
			fmt.Println(i, v)
		}

		// -------------------------------------

		for indice, value := range "AMOR"{
			fmt.Println(indice, string(value))
		}

		// -------------------------------------

		user := map[string]string {
			"Nome": "João",
			"Cidade": "Minas Gerais",
		}

		for chave, valor := range user {
			fmt.Println(chave, valor)
		}
}