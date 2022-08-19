package main

import ("fmt")

func fibonacci (p uint) uint {
	if p <= 1 {
		return p
	}

	return fibonacci(p - 2) + fibonacci(p - 1)
}

func main() {

	// 1 1 2 3 5 8 13
	p := uint(10)
	for i := uint(1); i < p; i++ {
		fmt.Println(fibonacci(i))
	}
	
}