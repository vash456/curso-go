package main

import "fmt"

func main() {
	// las variables dentro del if solo estan instanciasdas dentro del if
	if nombre, edad := "Alex", 26; nombre == "Alex" {
		fmt.Println("hola, ", nombre)
	} else {
		fmt.Printf("Nombre: %s - Edad %d\n", nombre, edad)
	}

	//obtener valor de mapa
	users := make(map[string]string)

	users["Alex"] = "alex@gmail.com"
	users["Roel"] = "roel@gmail.com"

	correo, ok := users["Alex"] //el ok es la validacion de que el valor existe
	fmt.Println(correo, ok)

	if correo, ok := users["juan"]; ok {
		fmt.Println(correo)
	} else {
		fmt.Println("No fue posible obtener el valor")
	}
}
