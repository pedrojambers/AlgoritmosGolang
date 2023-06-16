package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Por favor entre com um número")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	entrada := scanner.Text()

	numero, err := strconv.Atoi(entrada)
	if err != nil {
		fmt.Println("Erro ao ler o número:", err)
		return
	}

	if verificarPrimo(numero) {
		fmt.Println(numero, "é um número primo.")
	} else {
		fmt.Println(numero, "não é um número primo.")
	}
}

func verificarPrimo(numero int) bool {
	if numero < 2 {
		return false
	}

	for i := 2; i*i <= numero; i++ {
		if numero%i == 0 {
			return false
		}
	}

	return true
}