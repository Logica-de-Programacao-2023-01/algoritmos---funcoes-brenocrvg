package main

import (
	"errors"
	"fmt"
	"strings"
)

func extrair(texto string) ([]string, error) {
	if texto == "" {
		return nil, errors.New("string vazia")
	}

	palavras := strings.Fields(texto)
	return palavras, nil
}

func main() {
	frase := "olá galera legal, show"
	palavras, err := extrair(frase)
	if err != nil {
		fmt.Println("erro:", err)
	} else {
		fmt.Println("palavras:", palavras)
	}
}
