package main

import (
	"errors"
	"fmt"
	"math"
)

func primordio(numero int) (bool, error) {
	if numero < 0 {
		return false, errors.New("negativo")
	}

	if numero < 2 {
		return false, nil
	}

	limite := int(math.Sqrt(float64(numero)))
	for i := 2; i <= limite; i++ {
		if numero%i == 0 {
			return false, nil
		}
	}

	return true, nil
}

func main() {
	numero := 23
	primo, err := primordio(numero)
	if err != nil {
		fmt.Println("erro:", err)
	} else {
		fmt.Println("primo?", primo)
	}
}
