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
		fmt.Println("Hasta luego.")
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

	m1 := make(chan bool)
	var wg2 sync.WaitGroup

	wg2.Add(1)
	go VerMonitoreo(cpuData, abb, m1, &wg2, "Salir: s+Enter | Analizar proceso: a+Enter\n")

	go func() {
		var input1 string
		for {
			fmt.Scanln(&input1)
			if input1 == "s" {
				close(m1)
				wg2.Wait()
				close(done)
				break
			}
			if input1 == "a" {
				close(m1)
				wg2.Wait()

				m2 := make(chan bool)
				var wg3 sync.WaitGroup

				wg3.Add(1)

				go VerMonitoreo(cpuData, abb, m2, &wg3, "Analizador de procesos\nVer info: v+Enter | Rastrear: r+Enter | Dar de baja m+Enter")

				var input2 string
				fmt.Scanln(&input2)

				if input2 == "v" {
					close(m2)
					wg3.Wait()
					m3 := make(chan bool)
					var wg4 sync.WaitGroup

					wg4.Add(1)

					go VerMonitoreo(cpuData, abb, m3, &wg4, "Indique el PID+Enter para ver la info del proceso")

					var pid int
					fmt.Scanln(&pid)
					//podria ser igual que el rastrear
					close(m3)
					wg4.Wait()
					ui.Clear()
					monitor.VerProceso(pid)

					close(done)
					break

				}
				if input2 == "r" {
					close(m2)
					wg3.Wait()
					fmt.Println("Indique el PID+Enter")
					var pid int
					fmt.Scanln(&pid)
					ui.Clear()
					monitor.RastrearDetalleProceso(pid)
					close(done)
					break
				}
				if input2 == "m" {
					close(m2)
					wg3.Wait()
					fmt.Println("Indique el PID+Enter")
					var pid int
					fmt.Scanln(&pid)
					ui.Clear()
					monitor.DarDeBaja(pid)
					close(done)
					break
				}
			}
		}
	}()

	wg1.Wait()
	fmt.Println("Monitoreo terminado.")
}

func VerMonitoreo(cpuData <-chan float64, abb *utils.ABB, m <-chan bool, wg2 *sync.WaitGroup, opInfo string) {
	defer wg2.Done()

	for usage := range cpuData {
		select {
		case <-m:
			return
		default:
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
