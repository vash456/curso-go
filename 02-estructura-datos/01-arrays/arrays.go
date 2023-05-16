package main

import "fmt"

func main() {
	var numeros [5]int
	fmt.Println(numeros)

	numeros[0] = 10
	numeros[2] = 30

	fmt.Println(numeros)

	nombres := [3]string{"ab", "cd", "ef"}
	fmt.Println(nombres)

	colores := [...]string{
		"rojo",
		"verde",
		"azul",
		"negro",
	}

	fmt.Println(colores, len(colores))

	//indices indefinidos
	monedas := [...]string{
		0: "dolar",
		2: "bitcoin",
		3: "peso",
		5: "euro",
	}

	fmt.Println(monedas, len(monedas))

	//sub array
	subMoneda := monedas[0:3]
	fmt.Println(subMoneda, len(subMoneda))
}
