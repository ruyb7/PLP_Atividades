package main

import "fmt"

type Matematica struct{}

func Fatorial(n int) int {
	if n < 0 {
		return -1
	}
	if n == 0 || n == 1 {
		return 1
	}
	resultado := 1
	for i := 2; i <= n; i++ {
		resultado *= i
	}
	return resultado
}

func main() {
	fmt.Println(Fatorial(5))
}
