package main

import "fmt"

func main() {
	a := 2

	a = a + 2

	// Operadores asignacion

	a += 2
	fmt.Println(a)

	a -= 2
	fmt.Println(a)

	a *= 4
	fmt.Println(a)

	a /= 2
	fmt.Println(a)

	a %= 2
	fmt.Println(a)

	// Operador de incremento

	a = 0
	a++
	a++
	a++
	fmt.Println(a)

	// Operador de decremento

	a--
	a--
	a--
	fmt.Println(a)

}
