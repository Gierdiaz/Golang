package main

import ("fmt")

//Com isso você pode ignorar todos os tipos.
func generica (interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generica("string")
	generica(1)
	generica(true)

	mapa := map[interface{}]interface{} {
		1: "string",
		true: float64(15.54),
	}

	fmt.Println(mapa)
}