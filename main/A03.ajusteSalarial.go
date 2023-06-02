package main

import (
	"fmt"
)

func ajusteSalarial() {
	var salario float64
	var percentual float64
	var novoSalario float64

	fmt.Print("Digite o salário mensal: ")
	fmt.Scanln(&salario)

	fmt.Print("Digite o percentual de reajuste: ")
	fmt.Scanln(&percentual)

	novoSalario = salario + (salario * (percentual / 100.0))

	fmt.Printf("O novo salário após o reajuste é: R$ %.2f\n", novoSalario)
}
