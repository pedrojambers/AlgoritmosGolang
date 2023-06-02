package main

import "fmt"

func areaCircunferencia() {
	var raio float64

	fmt.Print("Digite o valor do raio da circunferencia: ")
	fmt.Scanln(&raio)

	fmt.Printf("Valor do raio: %.2f", (raio*raio)*3.14159265)
}
