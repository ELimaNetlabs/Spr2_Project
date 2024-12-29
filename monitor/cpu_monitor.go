package monitor

import (
	"Spr2_Project/utils" // Importa el paquete de ABB
	"fmt"
	"time"

	"sync"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/process"
)

// MonitoreoCPU se encarga de monitorear el uso de la CPU y los procesos más intensivos.
func MonitoreoCPU(data chan<- float64, wg *sync.WaitGroup, done <-chan bool, abb *utils.ABB) {
	defer wg.Done()

	for {
		select {
		case <-done:
			return
		default:
			// Obtener el uso total de CPU
			percentages, err := cpu.Percent(1*time.Second, false)
			if err != nil {
				fmt.Println("Error obteniendo uso de CPU:", err)
				continue
			}

			if len(percentages) > 0 {
				data <- percentages[0] // Enviar porcentaje al canal
			}

			// Obtener los procesos que más consumen CPU
			processes, err := process.Processes()
			if err != nil {
				fmt.Println("Error obteniendo procesos:", err)
				continue
			}

			for _, p := range processes {
				cpuUsage, err := p.CPUPercent()
				if err != nil {
					continue
				}

				// Obtener el nombre del proceso usando p.Name() en lugar de p.Exe
				name, err := p.Name()
				if err != nil {
					continue
				}

				// Crear un nodo con el proceso y su uso de CPU
				nodo := &utils.Nodo{
					Nombre:   name,
					PID:      int(p.Pid),
					CPUUsage: cpuUsage,
				}

				// Insertar el proceso en el ABB
				abb.Insertar(nodo)
			}
		}
	}
}
