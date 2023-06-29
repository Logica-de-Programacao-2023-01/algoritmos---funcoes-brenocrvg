package main

import (
	"errors"
	"fmt"
)

func inter(slice1, slice2 []int) ([]int, error) {
	if len(slice1) == 0 || len(slice2) == 0 {
		return nil, errors.New("um dos slices está vazio")
	}

	intersecao := []int{}
	mapa := make(map[int]bool)
	for _, num := range slice1 {
		mapa[num] = true
	}
	for _, num := range slice2 {
		if mapa[num] {
			intersecao = append(intersecao, num)
		}
	}

	return intersecao, nil
}

func main() {
	slice1 := []int{1, 2, 3, 4}
	slice2 := []int{3, 4, 5, 6}
	resultado, err := inter(slice1, slice2)
	if err != nil {
		fmt.Println("erro:", err)
	} else {
		fmt.Println("interseção:", resultado)
	}
}
