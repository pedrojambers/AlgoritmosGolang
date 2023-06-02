package main

import "fmt"

func somaQuadrados() {
	var v1, v2, v3 int

	fmt.Print("Digite o primeiro valor: ")
	fmt.Scan(&v1)

	fmt.Print("Digite o segundo valor: ")
	fmt.Scan(&v2)

	fmt.Print("Digite o terceiro valor: ")
	fmt.Scan(&v3)

	somaQuadrados := (v1 * v1) + (v2 * v2) + (v3 * v3)

	fmt.Printf("A soma dos quadrados dos tres valores e %d", somaQuadrados)
}
