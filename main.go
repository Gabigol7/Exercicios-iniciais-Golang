package main

import "fmt"

func main() {

	var nome string
	var idade int
	var peso int

	fmt.Print("Qual é o seu nome? ")
	fmt.Scan(&nome)
	fmt.Println("O seu nome é", nome)

	fmt.Print("Qual sua idade? ")
	fmt.Scan(&idade)
	fmt.Println("Sua idade é", idade)

	fmt.Print("Qual é o seu peso? ")
	fmt.Scan(&peso)
	fmt.Println("Seu peso é", peso)

}
