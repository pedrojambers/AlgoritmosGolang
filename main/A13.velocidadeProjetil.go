package main

import "fmt"

func velocidadeProjetil() {
	var velocidade float64
	var distancia float64
	var tempo float64

	fmt.Print("Informa a distancia percorrida em quilômetros ")
	fmt.Scan(&distancia)

	fmt.Print("Informa o tempo de percurso em minutos ")
	fmt.Scan(&tempo)

	velocidade = (distancia * 1000) / (tempo * 60)
	fmt.Printf("A velocidade do projétil, em metros por segundo, é %.2f", velocidade)
}
