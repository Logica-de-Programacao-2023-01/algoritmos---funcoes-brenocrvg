package main

import (
	"errors"
	"fmt"
	"sort"
)

func crescente(slice []int) ([]int, error) {
	if len(slice) == 0 {
		return nil, errors.New("O slice está vazio")
	}

	valores := make([]int, len(slice))
	copy(valores, slice)
	sort.Ints(valores)
	return valores, nil
}

func main() {
	numeros := []int{51435, 2643445, 806554, 16457868, 10000}
	valoresOrdenados, err := crescente(numeros)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("valores ordenados:", valoresOrdenados)
	}
}
