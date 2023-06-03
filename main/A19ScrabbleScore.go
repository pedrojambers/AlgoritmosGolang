package main

import (
	"fmt"
	"strings"
)

func main() {
	var palavra string

	fmt.Print("Digite uma palavra: ")
	fmt.Scanln(&palavra)

	palavra = strings.ToLower(palavra)

	score := 0

	for _, letra := range palavra {
		score += pontuacaoLetra(letra)
	}

	fmt.Printf("O score da palavra '%s' é %d\n", palavra, score)
}

func pontuacaoLetra(letra rune) int {
	switch letra {
	case 'a', 'e', 'i', 'o', 'u', 'l', 'n', 'r', 's', 't':
		return 1
	case 'd', 'g':
		return 2
	case 'b', 'c', 'm', 'p':
		return 3
	case 'f', 'h', 'v', 'w', 'y':
		return 4
	case 'k':
		return 5
	case 'j', 'x':
		return 8
	case 'q', 'z':
		return 10
	default:
		return 0
	}
}
