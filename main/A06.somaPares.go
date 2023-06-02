package main

import "fmt"

func somaPares() {
	var inicio, fim int

	fmt.Print("Digite o valor inicial do intervalo: ")
	fmt.Scanln(&inicio)

	fmt.Print("Digite o valor final do intervalo: ")
	fmt.Scanln(&fim)

	somaPares := 0

	for i := inicio; i <= fim; i++ {
		if i%2 == 0 {
			somaPares += i
		}
	}

	fmt.Printf("A soma dos números pares em um intervalo de %d a %d é: %d\n", inicio, fim, somaPares)
}
