package main

import (
	"fmt"
	"strings"
)

func cont(texto string) int {
	total := 0
	vogais := "aeiouAEIOU"

	for _, char := range texto {
		if strings.ContainsRune(vogais, char) {
			total++
		}
	}

	return total
}

func main() {
	texto := "Hello, World!"
	quantidadeVogais := cont(texto)
	fmt.Println("a quantidade de vogais é:", quantidadeVogais)
}
