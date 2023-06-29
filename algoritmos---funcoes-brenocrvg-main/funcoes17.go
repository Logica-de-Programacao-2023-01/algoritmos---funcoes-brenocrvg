package main

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

func c(slice []string) (string, error) {
	if len(slice) == 0 {
		return "", errors.New("vazio")
	}

	sort.Strings(slice)
	resultado := strings.Join(slice, "")

	return resultado, nil
}

func main() {
	slice := []string{"blusa", "palavra", "linguiça", "morango"}
	novaString, err := c(slice)
	if err != nil {
		fmt.Println("erro:", err)
	} else {
		fmt.Println("nova string:", novaString)
	}
}
