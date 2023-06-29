package main

import "fmt"

func encontrarPosicao(slice []int, valor int) int {
	for i, num := range slice {
		if num == valor {
			return i
		}
	}
	return -1
}

func main() {
	numeros := []int{232, 4123, 634, 678, 410}
	valor := 6
	posicao := encontrarPosicao(numeros, valor)
	fmt.Println("A posição do primeiro elemento igual a", valor, "é:", posicao)
}
