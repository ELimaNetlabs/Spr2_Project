package main

import (
	"Spr2_Project/monitor"
	"Spr2_Project/utils"

	//"bufio"
	"fmt"
	//"os"
	//"strings"
	"sync"
	"time"

	"github.com/spf13/cobra"
)

func main() {
	var cpuFlag bool
	var memFlag bool

	//REVISAR HERRAMIENTA
	var rootCmd = &cobra.Command{
		Use:   "monitoreo",
		Short: "Herramienta CLI para monitorear recursos del sistema",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Debe pasar un flag (--cpu || --mem)")
		},
	}

	rootCmd.PersistentFlags().BoolVar(&cpuFlag, "cpu", false, "Monitoreo de CPU")
	rootCmd.PersistentFlags().BoolVar(&memFlag, "mem", false, "Monitoreo de memoria")

	rootCmd.Execute()
	//

	if cpuFlag {
		monitorCPU()
	}
	if memFlag {
		monitorMemory()
	}
}

func monitorCPU() {
	fmt.Println("Monitoreando uso de CPU...")

	abb := &utils.ABB{}
	cpuData := make(chan float64)
	done := make(chan bool)
	var wg sync.WaitGroup

	wg.Add(1)
	go monitor.MonitoreoCPU(cpuData, &wg, done, abb)
	go func() {
		for usage := range cpuData {
			fmt.Printf("\r")
			// Limpiar la pantalla antes de imprimir (opcional)
			fmt.Printf("\nUso de CPU: %.2f%%", usage)
			// Mostrar los procesos almacenados en el ABB
			fmt.Println("\nProcesos con más uso de CPU (ordenados):")
			// Aquí podrías agregar una función para recorrer el ABB y mostrar los procesos
			abb.ListarTop5()
		}
	}()

	time.Sleep(10 * time.Second)
	close(done)
	wg.Wait()

	fmt.Println("Monitoreo terminado.")
}

func monitorMemory() {
	fmt.Println("Monitoreando uso de memoria...")
}
