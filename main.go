package main

import (
	"Spr2_Project/monitor" // Asegúrate de importar correctamente el paquete
	"Spr2_Project/utils"   // Paquete del ABB

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

	var rootCmd = &cobra.Command{
		Use:   "monitoreo",
		Short: "Herramienta CLI para monitorear recursos del sistema",
		Run: func(cmd *cobra.Command, args []string) {
			// Comportamiento por defecto si no se especifica un flag
			fmt.Println("Monitoreo en ejecución...")
		},
	}

	// Definir las flags
	rootCmd.PersistentFlags().BoolVar(&cpuFlag, "cpu", false, "Monitoreo de CPU")
	rootCmd.PersistentFlags().BoolVar(&memFlag, "mem", false, "Monitoreo de memoria")

	// Ejecutar el comando
	rootCmd.Execute()
	monitorCPU()
	// Lógica según los flags
	if cpuFlag {
		// Iniciar monitoreo de CPU
		monitorCPU()
	}
	if memFlag {
		// Iniciar monitoreo de memoria
		monitorMemory()
	}
}

func monitorCPU() {
	fmt.Println("Monitoreando uso de CPU...")
	// Crear el ABB
	abb := &utils.ABB{}

	// Crear canales y WaitGroup
	cpuData := make(chan float64)
	done := make(chan bool)
	var wg sync.WaitGroup

	// Agregar una goroutine al WaitGroup
	wg.Add(1)

	// Iniciar monitoreo de CPU y procesos
	go monitor.MonitoreoCPU(cpuData, &wg, done, abb)

	// Goroutine para procesar los datos de CPU
	go func() {
		for usage := range cpuData {
			// Limpiar la pantalla antes de imprimir (opcional)
			fmt.Printf("\rUso de CPU: %.2f%%", usage)
		}
	}()

	// Ejecutar monitoreo por 10 segundos
	time.Sleep(10 * time.Second)

	// Señalar que debe detenerse el monitoreo
	close(done)

	// Esperar a que las goroutines terminen
	wg.Wait()

	// Mostrar los procesos almacenados en el ABB
	fmt.Println("\nProcesos con más uso de CPU (ordenados):")
	// Aquí podrías agregar una función para recorrer el ABB y mostrar los procesos
	abb.ListarTop5()

	fmt.Println("Monitoreo terminado.")
}

func monitorMemory() {
	// Lógica para monitorear memoria
	fmt.Println("Monitoreando uso de memoria...")
}
