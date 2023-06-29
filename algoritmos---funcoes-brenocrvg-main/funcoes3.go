package main

import (
	"fmt"
	"strings"
)

func concatenar(slice []string) string {
	return strings.Join(slice, "")
}

func main() {
	strings := []string{"vasco", " da ", "gama"}
	concatenacao := concatenar(strings)
	fmt.Println("concatenação:", concatenacao)
}
