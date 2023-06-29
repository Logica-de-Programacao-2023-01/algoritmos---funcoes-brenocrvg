package main

import (
	"errors"
	"fmt"
)

func aplicarmassa(slice []int, funcao func(int) int) (int, error) {
	if len(slice) == 0 {
		return 0, errors.New("O slice está vazio")
	}

	resultado := 0
	for _, valor := range slice {
		resultado += funcao(valor)
	}

	return resultado, nil
}

func dobrarlegal(numero int) int {
	return numero * 2
}

func main() {
	slice := []int{16346, 263465, 345, 463678, 565677}
	soma, err := aplicarmassa(slice, dobrarlegal)
	if err != nil {
		fmt.Println("erro:", err)
	} else {
		fmt.Println("soma:", soma)
	}
}
