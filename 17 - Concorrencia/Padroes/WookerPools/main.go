package main

func fiboacci(n int) int {
	if n <= 1 {
		return n
	}

	return fiboacci(n-2) + fiboacci(n-1)
}

func main() {

	tarefas := make(chan int, 40)
	resultados := make(chan int, 40)

	// Criando 3 workers para processar as tarefas
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)

	for i := 0; 0 < 40; i++ {
		tarefas <- i
	}

	close(tarefas)

	for i := 0; i < 40; i++ {
		println(<-resultados)
	}

}

func worker(tarefas <-chan int, resultados chan<- int) {
	for n := range tarefas {
		resultados <- fiboacci(n)
	}
}
