package main

import "fmt"

func main() {
	fmt.Println("\nHOLA CONDICIONALES IF\n")

	valor1 := 1
	valor2 := 2

	if valor1 == 1 {
		fmt.Println("valor 1 es:", valor1, "\n")
	} else {
		fmt.Println("valor 1 es diferente de 1\n")
	}

	//operador AND: &&
	fmt.Println("	Operador AND: &&")
	if valor1 == 1 && valor2 == 2 {
		fmt.Println("Es valor1 = 1 y valor2 = 2 ??  -->  Es verdad\n")
	}

	//Operador OR: ||
	fmt.Println("	Operador OR: &&")
	if valor1 == 1 || valor2 == 3 {
		fmt.Println("Es valor1 = 1 ó valor2 = 2 ??  -->  Es verdad\n")
	}

	//Convertir texto a número
	//...se utiliza el paquete STRCONV, el cual retorna 2 valores: value y err
	value, err
}
