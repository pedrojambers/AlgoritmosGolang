package main

import (
	"fmt"
	"strings"
)

func main() {
	var frase string

	println(" Digite uma frase:")
	fmt.Scanln(&frase)

	if ehPangrama(frase) {
		println("A frase é um pangrama!")
	} else {
		println("A frase não é um pangrama. ")
	}
}

func ehPangrama(frase string) bool {
	frase = strings.ToLower(frase)

	letras := make(map[rune]bool)

	for _, char := range frase {
		if char >= 'a' && char <= 'z' {
			letras[char] = true
		}
	}

	return len(letras) == 26
}