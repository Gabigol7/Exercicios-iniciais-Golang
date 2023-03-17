package main

import "fmt"

func main() {

	var numero1 float64
	var numero2 float64
	var numero3 float64
	var numero4 float64

	fmt.Print("Qual valor do numero1: ")
	fmt.Scanln(&numero1)

	fmt.Print("Qual valor do numero2: ")
	fmt.Scanln(&numero2)

	fmt.Print("Qual valor do numero3: ")
	fmt.Scanln(&numero3)

	fmt.Print("Qual valor do numero4: ")
	fmt.Scanln(&numero4)

	calculo_da_media := (numero1 + numero2 + numero3 + numero4) / 4

	fmt.Println("A média aritmética é:", calculo_da_media)
}
