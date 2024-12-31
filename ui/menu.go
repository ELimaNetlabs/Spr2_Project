package ui

import (
	"fmt"
	"time"
)

func MainMenu() int {

	for {
		fmt.Print("\033[H\033[2J")
		fmt.Println("Programa de monitoreo para cpu y memoria ram")
		fmt.Println("Seleccione una opcion:")
		fmt.Println("1-Ver uso del cpu")
		fmt.Println("2-Ver uso de la memoria ram")

		var input int
		fmt.Scanln(&input)
		fmt.Println("Ingresaste:", input)

		switch input {
		case 1:
			fmt.Println("Elegiste la opción 1")
			return 1

		case 2:
			fmt.Println("Elegiste la opción 2")
			return 2
		default:
			fmt.Println("Opción no válida")
			time.Sleep(1 * time.Second)
		}

	}

}
