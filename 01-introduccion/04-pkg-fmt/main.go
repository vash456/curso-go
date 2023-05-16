package main

import "fmt"

func main() {
	hola := "hola"
	mundo := "mundo"

	fmt.Println(hola)
	fmt.Println(mundo)

	nombre := "Alex"
	edad := 26

	// print que se usa para dar formato %s-string %d-numero entero %v-cualquier tipo de dato
	fmt.Printf("Hola, %s y su edad es %d \n", nombre, edad)
	fmt.Printf("Hola, %v y su edad es %v \n", nombre, edad)

	mensaje := fmt.Sprintf("Hola, %s y su edad es %d", nombre, edad)
	fmt.Println(mensaje)

	fmt.Printf("variable tipo: %T \n", nombre)

	// recibir dato de consola
	fmt.Print("Ingrese otro nombre: ")
	fmt.Scanln(&nombre)

	fmt.Println("Otro nombre:", nombre)
}
