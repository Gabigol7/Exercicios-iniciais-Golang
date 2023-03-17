package main

import "fmt"

func main() {

	var base int
	var altura int

	fmt.Print("Valor da base? ")
	fmt.Scan(&base)
	fmt.Println("O resultado da base é", base)

	fmt.Print("Valor da altura? ")
	fmt.Scan(&altura)
	fmt.Println("O resultado da altura é", altura)

	fmt.Println((base * altura) / 2)

}
