package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("------ JoKenPo ------")
	fmt.Println("Jogador Um Escolha uma jogada:\n")
	fmt.Println("1. Pedra")
	fmt.Println("2. Papel")
	fmt.Println("3. Tesoura\n")

	var jogadorUm string
	var jogadorDois string
	opcaoValida := false

	for !opcaoValida {
		fmt.Print("Jogador Um: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		jogadorUm = scanner.Text()

		if jogadorUm == "1" || jogadorUm == "2" || jogadorUm == "3" {
			opcaoValida = true
		} else {
			fmt.Println("Um valor de 1 a 3!")
		}
	}

	fmt.Println("Jogador Dois Escolha uma jogada:\n")
	fmt.Println("1. Pedra")
	fmt.Println("2. Papel")
	fmt.Println("3. Tesoura")

	opcaoValida = false

	for !opcaoValida {
		fmt.Print("Jogador Dois: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		jogadorDois = scanner.Text()

		if jogadorDois == "1" || jogadorDois == "2" || jogadorDois == "3" {
			opcaoValida = true
		} else {
			fmt.Println("Um valor de 1 a 3!")
		}
	}

	resultado := ""

	if jogadorUm == jogadorDois {
		resultado = "Empate"
	} else if jogadorUm == "1" && jogadorDois == "3" {
		resultado = "Jogador Um venceu"
	} else if jogadorUm == "3" && jogadorDois == "2" {
		resultado = "Jogador Um venceu"
	} else if jogadorUm == "2" && jogadorDois == "1" {
		resultado = "Jogador Um venceu"
	} else {
		resultado = "Jogador Dois venceu"
	}

	fmt.Println("Resultado:", resultado)
}