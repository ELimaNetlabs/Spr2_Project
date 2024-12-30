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
	//Revisar herramienta
	var rootCmd = &cobra.Command{
		Use:   "monitoreo",
		Short: "Herramienta CLI para monitorear recursos del sistema",
		Run: func(cmd *cobra.Command, args []string) {
			if cpuFlag {
				monitorCPU()
			} else if memFlag {
				monitorMemory()
			} else {
				fmt.Println("Debe pasar un flag (--cpu || --mem)")
			}
		},
	}

	rootCmd.Flags().BoolVar(&cpuFlag, "cpu", false, "Monitoreo de CPU")
	rootCmd.Flags().BoolVar(&memFlag, "mem", false, "Monitoreo de memoria")

	rootCmd.Execute()
	//
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
			fmt.Printf("\nUso de CPU: %.2f%%", usage)
			fmt.Println("\nLos 5 procesos con más uso de CPU:")
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
