package main

import "fmt"

func main() {
	fmt.Println("Digite um nome:")
	var name string
	fmt.Scanln(&name)

	message := TwoFer(name)
	fmt.Println(message)
}

func TwoFer(name string) string {
	if name == "" {
		name = "você"
	}
	return fmt.Sprintf("Um para %s, um para mim.", name)
}
