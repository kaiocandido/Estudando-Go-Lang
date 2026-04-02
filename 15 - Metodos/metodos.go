package main

import "fmt"

// structs são tipos de dados compostos que agrupam variáveis relacionadas. Eles são usados para criar tipos personalizados que podem conter múltiplos campos. Por exemplo, podemos definir uma struct para representar uma pessoa:
type Pessoa struct {
	Nome  string
	Idade int
}

// Métodos são funções associadas a um tipo específico. Eles permitem que você defina comportamentos para os tipos personalizados que você cria. Em Go, os métodos são definidos usando a sintaxe de receptor, onde o receptor é o tipo ao qual o método pertence.
func (u Pessoa) Salvar() {
	fmt.Printf("Salvando a pessoa: %s, idade: %d\n", u.Nome, u.Idade)
}

// O método idade verifica se a pessoa é maior de idade (18 anos ou mais) e retorna um valor booleano indicando o resultado. Ele é definido como um método do tipo Pessoa, o que significa que pode ser chamado em qualquer instância da struct Pessoa.
func (u Pessoa) idade() bool {
	if u.Idade >= 18 {
		return true
	}
	return false
}

// O método aniversario é definido com um receptor de ponteiro (*Pessoa) para que ele possa modificar o valor da idade da pessoa. Ele incrementa a idade em 1, simulando o aniversário da pessoa.
func (u *Pessoa) aniversario() {
	u.Idade++
}

func main() {
	// Criando uma instância da struct Pessoa
	pessoa := Pessoa{
		Nome:  "João",
		Idade: 30,
	}

	fmt.Println(pessoa)
	pessoa.Salvar()

	idade := pessoa.idade()
	fmt.Println(idade)

	pessoa.aniversario()
	fmt.Println(pessoa.Idade)
}
