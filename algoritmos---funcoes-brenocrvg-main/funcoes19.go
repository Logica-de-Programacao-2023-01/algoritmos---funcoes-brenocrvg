package main

import (
	"errors"
	"fmt"
	"math"
)

func cousin(numero int) bool {
	if numero < 2 {
		return false
	}

	limite := int(math.Sqrt(float64(numero)))
	for i := 2; i <= limite; i++ {
		if numero%i == 0 {
			return false
		}
	}

	return true
}

func obter(numero int) ([]int, error) {
	if numero < 0 {
		return nil, errors.New("negativo")
	}

	primos := []int{}
	for i := 2; i <= numero; i++ {
		if cousin(i) {
			primos = append(primos, i)
		}
	}

	return primos, nil
}

func main() {
	numero := 20
	primos, err := obter(numero)
	if err != nil {
		fmt.Println("erro:", err)
	} else {
		fmt.Println("primos:", primos)
	}
}
