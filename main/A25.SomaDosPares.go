package main

import (
	"fmt"
)

func main() {
	var numeros []float64

	continuar := true

	for continuar {
		var numero float64


		println("| Digite um número:                   |")

		fmt.Scanln(&numero)

		numeros = append(numeros, numero)

		var opcao string

		println("| Deseja inserir mais números? (S/N): |")

		fmt.Scanln(&opcao)

		if opcao != "S" && opcao != "s" {
			continuar = false
		}
	}

	if len(numeros) > 0 {
		media := calcularMedia(numeros)

		fmt.Printf("A média dos números informados é: %.2f\n", media)

	} else {

		println("| Nenhum número foi informado.        |")

	}
}

func calcularMedia(numeros []float64) float64 {
	total := 0.0

	for _, numero := range numeros {
		total += numero
	}

	media := total / float64(len(numeros))
	return media
}