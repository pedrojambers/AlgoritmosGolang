package main

import (
	"fmt"
	"strings"
)

func main() {
	var entrada string

	fmt.Print("Digite uma palavra, frase ou número: ")
	fmt.Scanln(&entrada)

	entrada = strings.ToLower(entrada)
	if entrada == reverse(entrada) {
		fmt.Println("É um palíndromo")
	} else {
		fmt.Println("Não é um palíndromo")
	}
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
