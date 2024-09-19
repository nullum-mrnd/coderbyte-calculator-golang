package main

import "fmt"

func main() {
	var s string
	fmt.Print("Insira a expressão a ser calculada: ")
	fmt.Scanln(&s)
	fmt.Println("Resultado:", calculator(s))

}
