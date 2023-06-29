package main

import (
	"fmt"
	"sort"
)

func segundo(slice []int) (int, error) {
	tamanho := len(slice)
	if tamanho < 2 {
		return 0, fmt.Errorf("o slice precisa ter pelo menos dois elementos")
	}

	sort.Ints(slice)
	segundoMaior := slice[tamanho-2]
	return segundoMaior, nil
}

func main() {
	numeros := []int{324, 244, 277, 107, 82222}
	segundoMaior, err := segundo(numeros)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("o segundo maior valor é:", segundoMaior)
	}
}
