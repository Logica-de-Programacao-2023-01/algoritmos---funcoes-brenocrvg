package main

import (
	"errors"
	"fmt"
)

func presenca(numero int, slice []int) (bool, error) {
	if len(slice) == 0 {
		return false, errors.New("vazio")
	}

	for _, valor := range slice {
		if valor == numero {
			return true, nil
		}
	}

	return false, nil
}

func main() {
	numero := 5
	slice := []int{1, 2, 3, 4, 5}
	presente, err := presenca(numero, slice)
	if err != nil {
		fmt.Println("erro:", err)
	} else {
		fmt.Println("está presente?", presente)
	}
}
