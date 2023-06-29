package main

import (
	"errors"
	"fmt"
)

func strings5(slice []string) ([]string, error) {
	if len(slice) == 0 {
		return nil, errors.New("O slice está vazio")
	}

	resultado := []string{}
	for _, str := range slice {
		if len(str) > 5 {
			resultado = append(resultado, str)
		}
	}

	return resultado, nil
}

func main() {
	slice := []string{"show", "massa", "cobertura", "laddae", "aplicada"}
	novoSlice, err := strings5(slice)
	if err != nil {
		fmt.Println("erro:", err)
	} else {
		fmt.Println("novo slice:", novoSlice)
	}
}
