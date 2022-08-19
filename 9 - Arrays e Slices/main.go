package main

import ("fmt")

func main() {
	//--------------ARRAYS-------------
	array := [5]int{1,2,3,4,5}
	fmt.Println(array)

	array1 := [6]string{"A","B","C","D","E","F"}
	fmt.Println(array1)

	// ------------SLICES----------------
	slice := []int {1,2,3,4,5,6}
	
	slice = append(slice, 7)
	fmt.Println(slice)

	//inclusivo e exclusivo. O primeiro conta e o último sai [1:3]
	pieceSlice := array1[1:3]
	fmt.Println(pieceSlice)

	//alterando valor do array slice
	array1[1] = "Posição alterada"
	fmt.Println(pieceSlice)

	//------------ARRAYS INTERNOS----------------

	fmt.Println("======================")
	slice1 := make([]int, 3, 6)
	fmt.Println(slice1)
	fmt.Println(len(slice1)) //tamanho
	fmt.Println(cap(slice1)) //capacidade

	slice1 = append(slice1, 1,2,3,4,5)
	fmt.Println(slice1)

}