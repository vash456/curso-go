package main

import "fmt"

func main() {
	//slicen - parecidos a los arrays pero se puede cambiar la longitud

	numeros := []int{1, 2, 3}

	fmt.Println(numeros, len(numeros))

	numeros = append(numeros, 4)
	numeros = append(numeros, 5)

	fmt.Println(numeros, len(numeros))

	subSlicen := numeros[:2]

	numeros[0] = 100 // este cambio afecta al slicen subSlicen

	fmt.Println(subSlicen)
	fmt.Println(numeros)

	// caracteristicas slicen:
	// Punteros
	// longitud
	// Capacidad

	meses := []string{"enero", "febrero", "marzo"}

	fmt.Printf("Len: %v, Cap:%v, %p \n", len(meses), cap(meses), meses)

	meses = append(meses, "abril")

	fmt.Printf("Len: %v, Cap:%v, %p \n", len(meses), cap(meses), meses)
}
