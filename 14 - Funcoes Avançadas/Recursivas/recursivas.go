package main

func fiboacci(n int) int {
	if n <= 1 {
		return n
	}

	return fiboacci(n-2) + fiboacci(n-1)
}

func main() {

	posicao := 10

	for i := 0; i < posicao; i++ {
		println(fiboacci(i))
	}

}
