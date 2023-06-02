package main

import "fmt"

func volumeCaixa() {
	var (
		comprimento float64
		largura     float64
		altura      float64
	)

	fmt.Print("Comprimento: ")
	fmt.Scan(&comprimento)

	fmt.Print("Largura: ")
	fmt.Scan(&largura)

	fmt.Print("Altura: ")
	fmt.Scan(&altura)

	volume := comprimento * (largura * altura)

	fmt.Printf("O volume da caixa é: %.2f\n", volume)
}
