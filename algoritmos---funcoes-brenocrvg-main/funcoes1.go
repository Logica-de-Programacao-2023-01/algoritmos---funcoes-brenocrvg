package main

import "fmt"

func calcularmedia(slice []int) float64 {
	soma := 0
	tamanho := len(slice)
	for _, valor := range slice {
		soma += valor
	}
	media := float64(soma) / float64(tamanho)
	return media
}

func main() {
	numeros := []int{44, 12, 4, 991, 12234}
	media := calcularmedia(numeros)
	fmt.Println("média:", media)
}
