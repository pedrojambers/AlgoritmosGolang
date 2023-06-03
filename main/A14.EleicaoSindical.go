package main

import (
	"fmt"
)

func main() {
	var votosA, votosB, votosC, votosNulos, votosEmBranco int

	fmt.Print("Digite a quantidade de votos válidos para o candidato A: ")
	fmt.Scanln(&votosA)

	fmt.Print("Digite a quantidade de votos válidos para o candidato B: ")
	fmt.Scanln(&votosB)

	fmt.Print("Digite a quantidade de votos válidos para o candidato C: ")
	fmt.Scanln(&votosC)

	fmt.Print("Digite a quantidade de votos nulos: ")
	fmt.Scanln(&votosNulos)

	fmt.Print("Digite a quantidade de votos em branco: ")
	fmt.Scanln(&votosEmBranco)

	totalEleitores := votosA + votosB + votosC + votosNulos + votosEmBranco

	fmt.Println("---- Resultado da Eleição ----")
	fmt.Printf("Total de eleitores: %d\n", totalEleitores)
	fmt.Printf("Percentual de votos válidos: %.2f%%\n", float64(votosA+votosB+votosC)/float64(totalEleitores)*100)
	fmt.Printf("Percentual de votos válidos para o candidato A: %.2f%%\n", float64(votosA)/float64(totalEleitores)*100)
	fmt.Printf("Percentual de votos válidos para o candidato B: %.2f%%\n", float64(votosB)/float64(totalEleitores)*100)
	fmt.Printf("Percentual de votos válidos para o candidato C: %.2f%%\n", float64(votosC)/float64(totalEleitores)*100)
	fmt.Printf("Percentual de votos nulos: %.2f%%\n", float64(votosNulos)/float64(totalEleitores)*100)
	fmt.Printf("Percentual de votos em branco: %.2f%%\n", float64(votosEmBranco)/float64(totalEleitores)*100)
}
