package main

import (
	"fmt"
	"strings"
)

func reverse(cadena string) string {
	arrayCadena := strings.Split(cadena, "  ")
	fmt.Println(arrayCadena)
	arrayInvertida := make([]string, 0)
	for i := len(arrayCadena) - 1; i >= 0; i-- {
		arrayInvertida = append(arrayInvertida, arrayCadena[i])
	}
	return ""
}

func esPalindromo(palabra string) {
	fmt.Println(palabra)
	palabra = strings.ToLower(palabra)
	fmt.Println(palabra)
	//palabra = strings.ToUpper(palabra)
	//fmt.Println(palabra)
	palabra = strings.Replace(palabra, "z", "s", 2)
	fmt.Println(palabra)
}

func main() {
	esPalindromo("luz azul")
}
