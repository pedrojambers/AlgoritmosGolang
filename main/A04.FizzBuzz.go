package main

import (
	"fmt"
)

func fizzBuzz(start, end int) {
	if end <= start {
		fmt.Println("O número final deve ser maior que o número inicial.")
		return
	}

	for i := start; i <= end; i++ {
		output := ""

		if i%3 == 0 {
			output += "Fizz"
		}
		if i%5 == 0 {
			output += "Buzz"
		}

		if output == "" {
			output = fmt.Sprint(i)
		}

		fmt.Println(output)
	}
}

func main() {
	var start, end int

	fmt.Print("Digite o número inicial do intervalo: ")
	fmt.Scanln(&start)

	fmt.Print("Digite o número final do intervalo: ")
	fmt.Scanln(&end)

	fizzBuzz(start, end)
}
