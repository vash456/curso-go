package main

import "fmt"

func main() {
	dias := make(map[int]string)

	fmt.Println(dias)

	dias[1] = "domingo"
	dias[2] = "lunes"
	dias[3] = "martes"
	dias[4] = "miercoles"
	dias[5] = "jueves"
	dias[6] = "viernes"
	dias[7] = "sabado"

	fmt.Println(dias)

	delete(dias, 1)

	fmt.Println(dias, len(dias))

	// nuevo mapa
	estudiantes := make(map[string][]int)

	estudiantes["alex"] = []int{12, 45, 78}
	estudiantes["juan"] = []int{95, 32, 44}

	fmt.Println(estudiantes)

	fmt.Println(estudiantes["alex"][1])
}
