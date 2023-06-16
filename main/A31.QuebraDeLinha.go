package main

import (
	"fmt"
	"strings"
)

func main() {
	var frase string
	var colunas int

	fmt.Println("Digite a frase: ")
	_, err := fmt.Scan(&frase)
	if err != nil {
		return
	}

	fmt.Println("Digite a quantidade de colunas: ")
	_, err = fmt.Scan(&colunas)
	if err != nil {
		return
	}

	palavras := strings.Split(frase, " ")
	linhas := make([]string, 0)
	linhaAtual := ""
	for _, palavra := range palavras {
		if len(linhaAtual)+len(palavra)+1 <= colunas {
			linhaAtual += palavra + " "
		} else {
			linhas = append(linhas, linhaAtual)
			linhaAtual = palavra + " "
		}
	}
	linhas = append(linhas, linhaAtual)

	fmt.Println("Resultado:")
	for _, linha := range linhas {
		fmt.Println(linha)
	}
}