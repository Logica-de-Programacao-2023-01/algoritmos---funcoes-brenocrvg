package main

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

func filtrar(slice []string) (string, error) {
	if len(slice) == 0 {
		return "", errors.New("O slice está vazio")
	}

	var result strings.Builder
	for _, str := range slice {
		if len(str) > 0 && unicode.IsUpper(rune(str[0])) {
			result.WriteString(str)
			result.WriteString(" ")
		}
	}

	return strings.TrimSpace(result.String()), nil
}

func main() {
	palavras := []string{"vasco", "gigante", "enorme", "colossal"}
	resultado, err := filtrar(palavras)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Resultado:", resultado)
	}
}
