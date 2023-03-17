package main

import "fmt"

func main() {

	var real float64

	fmt.Print("Qual valor em dolar ?")
	fmt.Scan(&real)
	fmt.Println("O seu valor em real é:", real*5.20)

	//var dolar float64
	//fmt.Print("Qual valor do dolar ?")
	//fmt.Scan(&dolar)

	//conversão_dolar_para_real := (dolar * 20)

	//fmt.Println("A conversão do dolar é:", conversão_dolar_para_real)
}
