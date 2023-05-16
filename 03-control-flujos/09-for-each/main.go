package main

import "fmt"

func main() {
	nombres := [...]string{"alex", "roel", "juan"}

	// for-each
	for indice, elemento := range nombres {
		fmt.Println(indice, elemento)
	}

	for _, elemento := range nombres { // la barra baja es por si no se necesita la vriable indices
		fmt.Println(elemento)
	}

}
