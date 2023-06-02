package main

import "fmt"

func troco() {
	var total, pago float64

	fmt.Print("Digite o valor total a ser pago: ")
	fmt.Scanln(&total)

	fmt.Print("Digite o valor efetivamente pago: ")
	fmt.Scanln(&pago)

	troco := pago - total

	if troco < 0 {
		fmt.Println("Valor pago insuficiente.")
		return
	}

	cedulas := []float64{100, 50, 10, 5, 1}
	moedas := []float64{0.5, 0.1, 0.05, 0.01}

	fmt.Println("Troco:")

	for _, cedula := range cedulas {
		qtdCedulas := int(troco / cedula)
		if qtdCedulas > 0 {
			fmt.Printf("%d cédula(s) de R$%.2f\n", qtdCedulas, cedula)
			troco -= float64(qtdCedulas) * cedula
		}
	}

	for _, moeda := range moedas {
		qtdMoedas := int(troco / moeda)
		if qtdMoedas > 0 {
			fmt.Printf("%d moeda(s) de R$%.2f\n", qtdMoedas, moeda)
			troco -= float64(qtdMoedas) * moeda
		}
	}
}
