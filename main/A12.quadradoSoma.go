package main

import "fmt"

func quadradoSoma() {
	var v1, v2, v3 int

	fmt.Print("Digite o primeiro valor: ")
	fmt.Scan(&v1)

	fmt.Print("Digite o segundo valor: ")
	fmt.Scan(&v2)

	fmt.Print("Digite o terceiro valor: ")
	fmt.Scan(&v3)

	quadradoSoma := (v1 + v2 + v3) * (v1 + v2 + v3)

	fmt.Printf("o quadrado da soma dos tres valores e %d", quadradoSoma)
}
