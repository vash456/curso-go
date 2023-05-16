package main

import "fmt"

func main() {
	// Bucle tipo while: no existe while en go
	numeros := 12455
	c := 0
	for numeros > 0 {
		numeros /= 10
		c++
	}

	fmt.Println("Cantididad de digitos es: ", c)

	//  Bucle for
	for i := 0; i < 100; i++ {
		if i%2 == 0 { //imprimir solo pares
			fmt.Println(i)
		}
	}

	// break
	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Println("salta iteracion")
			break
		}
	}
	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Println("salta iteracion")
			continue
		}
	}
}
