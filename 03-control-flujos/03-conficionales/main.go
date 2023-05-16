package main

import "fmt"

func main() {
	var consumo, descuento float64
	var datosDescuento string

	fmt.Print("ingrese consumo: ")
	fmt.Scanln(&consumo)

	if consumo >= 0 {
		datosDescuento = "10%"
		descuento = consumo * 0.10
		fmt.Println()
	} else {
		fmt.Println("Error al ingresar")
	}

	montoDescuento := consumo - descuento
	igv := montoDescuento * 0.19
	totalPagar := montoDescuento + igv

	fmt.Println("------ FACTURA DE CONSUMO ------")
	fmt.Println("Descuento que aplica: ", datosDescuento)
	fmt.Println("Consumo: ", consumo)
	fmt.Println("Descuento: ", descuento)
	fmt.Println("Monto con descuento: ", montoDescuento)
	fmt.Println("IGV: ", igv)
	fmt.Println("Total a pagar: ", totalPagar)

}
