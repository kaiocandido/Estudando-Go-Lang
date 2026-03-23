package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int16 = 100 //int & 8 -> mostra a quantidade de bits que o tipo de dado tem, ou seja, a quantidade de números que ele pode armazenar
	//int8,int16,int32,int64
	fmt.Println(numero)

	var numero2 uint = 1000 //unsygned int -> só armazena números positivos, ou seja, o dobro do valor máximo do int
	//uint8,uint16,uint32,uint64
	fmt.Println(numero2)

	// alias INT32
	var numero3 rune = 1234
	fmt.Println(numero3)

	// alias UINT8
	var numero4 byte = 255
	fmt.Println(numero4)

	var float32, float64 = 3.14, 3.14159265358979323846264338327950288419716939937510 //float32 -> 6 casas decimais, float64 -> 15 casas decimais
	fmt.Println(float32, "|", float64)

	var str string = "Olá, mundo!" //string -> sequência de caracteres
	fmt.Println(str)

	char := 'A' //char -> tipo de dado que armazena um único caractere, representado por aspas simples
	fmt.Println(char)

	var booleano bool = true //bool -> tipo de dado que armazena valores booleanos (true ou false)
	fmt.Println(booleano)

	var erro error = errors.New("erro ocorrido") //error -> tipo de dado que representa um erro, utilizado para tratamento de erros
	fmt.Println(erro)

}
