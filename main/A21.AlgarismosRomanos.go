package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Insira uma data em Algarismos romanos e irei converter para número inteiro!")
	fmt.Print("Algarismo Romano: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	conversao := conversaoRomanoParaInt(input)

	fmt.Println("Em número inteiro fica:", conversao)
}

func conversaoRomanoParaInt(valor string) int {
	resultado := 0
	i := 0

	for i < len(valor) {
		charCorrespondente := valoresRomanos(valor[i])
		if i < len(valor)-1 {
			proximoChar := valoresRomanos(valor[i+1])
			if proximoChar > charCorrespondente {
				resultado += proximoChar - charCorrespondente
				i += 2
				continue
			}
		}
		resultado += charCorrespondente
		i++
	}

	return resultado
}

func valoresRomanos(c byte) int {
	switch c {
	case 'I', 'i':
		return 1
	case 'V', 'v':
		return 5
	case 'X', 'x':
		return 10
	case 'L', 'l':
		return 50
	case 'C', 'c':
		return 100
	case 'D', 'd':
		return 500
	case 'M', 'm':
		return 1000
	default:
		panic(fmt.Sprintf("Caractere inválido: %c", c))
	}
}