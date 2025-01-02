package ui

import (
	"fmt"
	"time"
)

func MainMenu() int {

	for {
		Clear()
		fmt.Println("Programa de monitoreo para cpu")
		fmt.Println("Iniciar? s/n")

		var input string
		fmt.Scanln(&input)

		switch input {
		case "s":
			return 1

		case "n":
			return 2
		default:
			fmt.Println("Opción no válida")
			time.Sleep(2 * time.Second)
		}

	}

}
