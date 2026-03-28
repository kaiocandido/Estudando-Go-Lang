package main

import "fmt"

func diaDaSemana(dia int) string {
	switch dia {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-feira"
	case 3:
		return "Terça-feira"
	case 4:
		return "Quarta-feira"
	case 5:
		return "Quinta-feira"
	case 6:
		return "Sexta-feira"
	case 7:
		return "Sábado"
	default:
		return "Dia inválido"
	}
}

func diaDaSemana2(dia int) string {
	switch {
	case dia == 1:
		return "Domingo"
	case dia == 2:
		return "Segunda-feira"
	case dia == 3:
		return "Terça-feira"
	case dia == 4:
		return "Quarta-feira"
	case dia == 5:
		return "Quinta-feira"
	case dia == 6:
		return "Sexta-feira"
	case dia == 7:
		return "Sábado"
	default:
		return "Dia inválido"
	}
}

func main() {

	fmt.Println(diaDaSemana(1))
	fmt.Println(diaDaSemana(3))
	fmt.Println(diaDaSemana(5))
	fmt.Println(diaDaSemana(7))
	fmt.Println(diaDaSemana(8))

}
