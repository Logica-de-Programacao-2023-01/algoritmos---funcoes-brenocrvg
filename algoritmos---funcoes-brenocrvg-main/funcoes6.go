package main

import (
	"errors"
	"fmt"
	"strings"
)

func conca(slice []string) (string, error) {
	if len(slice) == 0 {
		return "", errors.New("O slice está vazio")
	}

	concatenado := strings.Join(slice, ", ")
	return concatenado, nil
}

func main() {
	strings := []string{"tamo", "junto", "brincadeira"}
	concatenacao, err := conca(strings)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("concatenação com vírgulas:", concatenacao)
	}
}
