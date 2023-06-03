package main

import "fmt"

func main() {
	inicio := 1
	fim := 100

	fmt.Printf("Intervalo: %d - %d\n", inicio, fim)

	soma := 0
	for i := inicio; i <= fim; i++ {
		soma += i
	}

	fmt.Println("Soma:", soma)
}
