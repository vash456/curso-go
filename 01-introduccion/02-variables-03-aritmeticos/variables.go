package main

import "fmt"

func main() {
	var nombre1 string // variable estatica
	nombre1 = "Darlin"

	edad := 26 // variable dinamica

	fmt.Println(nombre1, edad)

	var a int     // 0
	var b float64 // 0
	var c string  // NULL
	var d bool    // false

	fmt.Println(a, b, c, d)

	const pi = 3.1415 // constante

	fmt.Println(pi)

	x := 20
	y := 10

	result := x + y
	result = x - y
	result = x * y
	result = x / y
	result = x % y

	fmt.Println("Resultado:", result)

	var div float64 = 3.0 / 2.0
	fmt.Println("Resultado:", div)

}
