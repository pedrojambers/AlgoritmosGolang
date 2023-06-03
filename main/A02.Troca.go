package main

import "fmt"

func main() {
	var A, B int
	fmt.Println("Digite um Valor para A:")
	fmt.Scanln(&A)
	fmt.Println("Digite um Valor para B:")
	fmt.Scanln(&B)

	valor := A
	A = B
	B = valor

	fmt.Println("Valores após serem trocados:")
	fmt.Println("A:", A)
	fmt.Println("A:", B)

}

//A2. Troca
//➢ Crie um programa que leia dois valores para as variáveis A e B
//➢ Efetue a troca dos valores de forma que a variável A passe a possuir o valor da variável B
//➢ A variável B passe a possuir o valor da variável A.
//➢ Apresente os valores após a efetivação do processamento da troca.
