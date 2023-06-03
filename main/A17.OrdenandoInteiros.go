package main

import (
	"fmt"
	"sort"
)

func main() {
	numeros := make([]int, 12)

	for i := 0; i < 12; i++ {
		fmt.Printf("Digite o elemento %d: ", i+1)
		fmt.Scanln(&numeros[i])
	}

	sort.Slice(numeros, func(i, j int) bool {
		return numeros[i] > numeros[j]
	})

	fmt.Println("Elementos em ordem decrescente:")
	fmt.Println(numeros)
}
