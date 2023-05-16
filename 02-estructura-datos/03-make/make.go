package main

import "fmt"

func main() {

	numeros := make([]int, 2, 3) // define slicen con 2 longitud 3 capacidad

	numeros[0] = 100
	numeros[1] = 200

	numeros = append(numeros, 400)

	fmt.Println(numeros, len(numeros), cap(numeros))

}
