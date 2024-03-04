package main

import "fmt"

func main() {
	saludar()
	saludoPersonalizado("hola Jupiter!!")
	tripleArgumento(3, 98, "jaja")

	//Ejemplo uso de una función que retorna un valor
	tripleArgumento(retornarValorEntero(21), 3, "hola!!")
	value := retornarValorEntero(33)
	fmt.Println("el valor retornado en la función es: ", value)

	//Ejemplo uso de una función que retorna Dos valores
	value1, value2 := dobleRetorno(4)
	fmt.Println(
		"\nEl valor de value1 es:", value1,
		"\nEl valor de value2 es:", value2,
	)

	//Ejemplo uso de una función que retorna Dos valores pero
	//solo se usa uno
	value3, _ := dobleRetorno(4)
	fmt.Println(
		"\nEl valor de value3 es:", value3,
	)
	_, value4 := dobleRetorno(4)
	fmt.Println(
		"El valor de value4 es:", value4,
	)

}

//********************************************************************

//Función básica sin argumentos
func saludar() {
	fmt.Println("\nHola Goland!!\n")
}

//Función básica con 1 argumento
func saludoPersonalizado(message string) {
	fmt.Println(message)
}

//Función básica con 3 argumentos (2 formas de declarar)
//func tripleArgumento(a int, b int, c string) {
func tripleArgumento(a, b int, c string) {
	fmt.Println(
		"\nargumento 1: ", a,
		"\nargumento 2: ", b,
		"\nargumento 3: ", c,
	)
}

//Función Retornando int
func retornarValorEntero(a int) int {
	return a * 2
}

//Función Retornando 2 valores
func dobleRetorno(a int) (c, d int) {
	return a, a * 2
}
