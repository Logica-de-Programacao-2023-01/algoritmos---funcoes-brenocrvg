package main

import (
	"errors"
	"fmt"
)

func filtro(slice []int) ([]int, error) {
	if len(slice) == 0 {
		return nil, errors.New("slice vazio")
	}

	pares := make([]int, 0)
	for _, valor := range slice {
		if valor%2 == 0 {
			pares = append(pares, valor)
		}
	}

	return pares, nil
}

func main() {
	numeros := []int{1, 2, 3, 4, 5, 6}
	pares, err := filtro(numeros)
	if err != nil {
		fmt.Println("erro:", err)
	} else {
		fmt.Println("pares:", pares)
	}
}
