package main

import "fmt"

func conversaoDolar() {
	var (
		cotacao float64
		reais   float64
	)

	fmt.Print("Valor da cotacao do dolar: ")
	fmt.Scan(&cotacao)

	fmt.Print("Quantidade de reais: ")
	fmt.Scan(&reais)

	conversao := reais / cotacao

	fmt.Printf("Os reais equivalem a U$ %.2f", conversao)
}
