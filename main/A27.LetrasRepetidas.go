package main

import (
	"fmt"
	"strings"
)

func main() {
	var texto string

	println("| Digite o texto:  |")
	fmt.Scanln(&texto)

	primeiraLetraNaoRepetida := encontrarPrimeiraLetraNaoRepetida(texto)

	if primeiraLetraNaoRepetida != "" {
		fmt.Printf("A primeira letra não repetida no texto é: %s\n", primeiraLetraNaoRepetida)
	} else {
		println("| Não existem letras que não se repetem no texto informado. |")
	}
}

func encontrarPrimeiraLetraNaoRepetida(texto string) string {
	letras := make(map[string]int)

	texto = strings.ToLower(texto)

	for _, letra := range texto {
		letras[string(letra)]++
	}

	for _, letra := range texto {
		if letras[string(letra)] == 1 {
			return string(letra)
		}
	}

	return ""
}