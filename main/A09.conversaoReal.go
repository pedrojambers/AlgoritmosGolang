package main

import "fmt"

func conversaoReal() {
	var (
		cotacao float64
		dolares float64
	)

	fmt.Print("Valor da cotacao do dolar: ")
	fmt.Scan(&cotacao)

	fmt.Print("Quantidade de dolares: ")
	fmt.Scan(&dolares)

	conversao := cotacao * dolares

	fmt.Printf("Os dolares equivalem a R$ %.2f", conversao)
}
