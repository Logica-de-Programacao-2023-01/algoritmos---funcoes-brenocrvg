package main

import (
	"errors"
	"fmt"
)

func aplicar(slice []int, funcao func(int) int) ([]int, error) {
	if len(slice) == 0 {
		return nil, errors.New("O slice está vazio")
	}

	resultados := make([]int, len(slice))
	for i, valor := range slice {
		resultados[i] = funcao(valor)
	}

	return resultados, nil
}

func dobrar(valor int) int {
	return valor * 2
}

func main() {
	numeros := []int{109, 487, 309, 542, 222}
	resultados, err := aplicar(numeros, dobrar)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Resultados:", resultados)
	}
}
