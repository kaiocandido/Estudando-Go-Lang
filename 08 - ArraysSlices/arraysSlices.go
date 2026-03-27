package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays and Slices")

	//PRIMEIRA FORMA DE CRIAR UM ARRAY
	var array1 [5]int
	array1[0] = 10
	array1[1] = 20
	array1[2] = 30
	array1[3] = 40
	array1[4] = 50
	fmt.Println("Array 1:", array1)

	//SEGUNDA FORMA DE CRIAR UM ARRAY
	array2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array 2:", array2)

	//TERCEIRA FORMA DE CRIAR UM ARRAY
	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println("Array 3:", array3)

	//CRIANDO UM SLICE

	slice := []int{10, 20, 30, 40, 50}
	fmt.Println("Slice:", slice)
	slice = append(slice, 60)
	fmt.Println("Slice após append:", slice)

	// CRIANDO UM SLICE A PARTIR DE UM ARRAY
	slice2 := array1[1:4] // Acessa os elementos do índice 1 ao 3 (exclusivo o índice 4)
	fmt.Println("Slice 2:", slice2)

	//TIPOS DE DADOS
	fmt.Println(reflect.TypeOf(array1))
	fmt.Println(reflect.TypeOf(slice))

	//Arrays internos

	slice3 := make([]int, 5, 10) // Cria um slice com capacidade para 5 elementos
	fmt.Println("Slice 3:", slice3)
	fmt.Println("Length:", len(slice3))   // Imprime o comprimento do slice
	fmt.Println("Capacity:", cap(slice3)) // Imprime a capacidade do slice

}
