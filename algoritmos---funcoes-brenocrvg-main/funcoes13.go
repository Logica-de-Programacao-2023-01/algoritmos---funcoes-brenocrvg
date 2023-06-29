package main

import (
	"errors"
	"fmt"
	"strconv"
)

func somad(numero int) (int, error) {
	if numero < 0 {
		return 0, errors.New("negativo")
	}

	soma := 0
	numeroStr := strconv.Itoa(numero)
	for _, char := range numeroStr {
		digito, _ := strconv.Atoi(string(char))
		soma += digito
	}

	return soma, nil
}

func main() {
	numero := 12345
	soma, err := somad(numero)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Soma dos dígitos:", soma)
	}
}
