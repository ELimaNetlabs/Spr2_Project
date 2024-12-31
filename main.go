package main

import (
	"Spr2_Project/monitor"
	"Spr2_Project/ui"
	"Spr2_Project/utils"
	"fmt"

	//"os"
	//"strings"
	"sync"
)

func main() {
	op := ui.MainMenu()

	if op == 1 {
		fmt.Print("\033[H\033[2J")
		monitorCPU()
	}
	if op == 2 {
		fmt.Print("\033[H\033[2J")
		monitorMemory()
	}

}

func monitorCPU() {

	abb := &utils.ABB{}
	cpuData := make(chan float64)
	done := make(chan bool)
	var wg sync.WaitGroup

	wg.Add(1)
	go monitor.MonitoreoCPU(cpuData, &wg, done, abb)

	go func() {
		fmt.Print("\033[H\033[2J")
		fmt.Println("Salir: s+Enter | Pausar: p+Enter")
		var input1 string
		for {
			fmt.Scanln(&input1)
			if input1 == "s" {
				close(done)
				break
			}
			if input1 == "p" {
				close(done)
				fmt.Println("Monitoreo pausado")
				fmt.Println("Ver info: v+Enter | Rastrear: r+Enter")
				var input2 string
				fmt.Scanln(&input2)
				if input2 == "v" {
					go monitor.VerProceso(input2)
					break
				}
				if input2 == "r" {
					go monitor.VerProceso(input2)
					break
				}
			}
		}
	}()

	go func() {
		for usage := range cpuData {
			fmt.Print("\033[2;1H")
			fmt.Println("Monitoreando uso de CPU...")
			fmt.Printf("\nUso de CPU: %.2f%%", usage)
			fmt.Println("\nLos 5 procesos con más uso de CPU:")
			abb.ListarTop5()
		}
	}()

	wg.Wait()
	fmt.Println("Monitoreo terminado.")
}

func monitorMemory() {
	fmt.Println("Monitoreando uso de memoria...")
}
