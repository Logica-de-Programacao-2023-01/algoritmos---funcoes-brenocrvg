package main

import (
	"errors"
	"fmt"
	"strings"
)

func substituir(str string) (string, error) {
	if str == "" {
		return "", errors.New("A string está vazia")
	}

	str = strings.ToLower(str)
	vogais := "aeiou"
	resultado := ""
	for _, char := range str {
		if strings.ContainsRune(vogais, char) {
			resultado += "*"
		} else {
			resultado += string(char)
		}
	}

	return resultado, nil
}

func main() {
	str := "vasco da gama tantas vezes campeao"
	novaStr, err := substituir(str)
	if err != nil {
		fmt.Println("erro:", err)
	} else {
		fmt.Println("nova string:", novaStr)
	}
}
