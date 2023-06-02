package main

import "fmt"

func salarioProfessor() {
	var hrsTrabalhadas float64
	var valorHora float64
	var inss float64

	fmt.Print("Numero de horas trabalhadas no mes: ")
	fmt.Scanln(&hrsTrabalhadas)

	fmt.Print("Valor hora-aula: ")
	fmt.Scanln(&valorHora)

	fmt.Print("Percentual de desconto do INSS: ")
	fmt.Scanln(&inss)

	salarioBruto := hrsTrabalhadas * valorHora
	descontoTotal := salarioBruto * (inss / 100)
	salarioLiquido := salarioBruto - descontoTotal

	fmt.Printf("Salario bruto: %.2f\n", salarioBruto)
	fmt.Printf("Salario líquido: %.2f", salarioLiquido)
}
