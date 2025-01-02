package main

import (
	"Spr2_Project/monitor"
	"Spr2_Project/ui"
	"Spr2_Project/utils"
	"fmt"
	"sync"
)

func main() {
	op := ui.MainMenu()

	switch op {
	case 1:
		ui.Clear()
		monitorCPU()
	case 2:
		ui.Clear()
		monitorMemory()
	default:
		fmt.Println("Algo salio mal...")
	}
}

func monitorCPU() {

	abb := &utils.ABB{}
	cpuData := make(chan float64)
	done := make(chan bool)
	var wg1 sync.WaitGroup

	wg1.Add(1)
	go monitor.MonitoreoCPU(cpuData, &wg1, done, abb)

	go func() {
		m1 := make(chan bool)
		go VerMonitoreo(cpuData, abb, m1, "Salir: s+Enter | Analizar proceso: a+Enter\n")

		var input1 string
		for {
			fmt.Scanln(&input1)
			if input1 == "s" {
				close(done)
				break
			}
			if input1 == "a" {
				close(m1)
				m2 := make(chan bool)
				VerMonitoreo(cpuData, abb, m2, "Analizador de procesos\nVer info: v+Enter | Rastrear: r+Enter")
				var input2 string
				fmt.Scanln(&input2)
				if input2 == "v" {
					ui.Clear()
					fmt.Println("Indique el PID+Enter para ver la info del proceso")

					var pid int
					fmt.Scanln(&pid)

					pInfo := make(chan string)
					var wg2 sync.WaitGroup
					wg2.Add(1)
					go monitor.VerProceso(pInfo, &wg2, pid)
					go func() {
						ui.Clear()
						fmt.Println("Debajo del top 5 esta la info del proceso elegido")
						for data := range pInfo {
							fmt.Println(data)
						}
					}()
					wg2.Wait()

					break
				}
				if input2 == "r" {
					done := make(chan bool)
					go func() {
						var input string
						fmt.Println("Escribe 's' para detener el rastreo.")
						fmt.Scanln(&input)
						if input == "s" {
							close(done)
						}
					}()
					var pid int
					fmt.Scanln(&pid)
					monitor.RastrearDetalleProceso(pid, done)
					break
				}
			}
		}
	}()

	wg1.Wait()
	fmt.Println("Monitoreo terminado.")
}

func VerMonitoreo(cpuData <-chan float64, abb *utils.ABB, m <-chan bool, opInfo string) {
	select {
	case <-m:
		return
	default:
		for usage := range cpuData {
			ui.Clear()
			fmt.Println(opInfo)
			fmt.Print("Monitoreando uso de CPU...\n")
			fmt.Printf("Uso de CPU: %.2f%%", usage)

			if abb.Raiz == nil {
				fmt.Println("\nEsperando datos...")
				continue
			}

			fmt.Println("\nLos 5 procesos con más uso de CPU:")
			abb.ListarTop5()
		}
	}
}

func monitorMemory() {
	fmt.Println("Monitoreando uso de memoria...")
}
