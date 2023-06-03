package main

import "fmt"

func main() {
	var nota1, nota2, nota3, nota4 float64

	fmt.Print("Digite a nota do 1º bimestre: ")
	fmt.Scanln(&nota1)

	fmt.Print("Digite a nota do 2º bimestre: ")
	fmt.Scanln(&nota2)

	fmt.Print("Digite a nota do 3º bimestre: ")
	fmt.Scanln(&nota3)

	fmt.Print("Digite a nota do 4º bimestre: ")
	fmt.Scanln(&nota4)

	media := (nota1 + nota2 + nota3 + nota4) / 4

	fmt.Printf("Média do aluno: %.2f\n", media)

	if media >= 5 {
		fmt.Println("Aprovado")
	} else {
		fmt.Println("Reprovado")
	}
}
