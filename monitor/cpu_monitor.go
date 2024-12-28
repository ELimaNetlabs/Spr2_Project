package monitor

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
)

// MonitorCPU inicia un monitoreo continuo del uso de la CPU.
// Los datos se envían a través del canal cpuData.
func MonitorCPU(cpuData chan<- float64, done <-chan bool) {
	for {
		select {
		case <-done: // Finalizar monitoreo si se recibe una señal en el canal done
			fmt.Println("Monitoreo de CPU detenido.")
			return
		default:
			// Obtiene el porcentaje de uso de la CPU
			percentages, err := cpu.Percent(1*time.Second, false)
			if err != nil {
				fmt.Println("Error obteniendo datos de CPU:", err)
				continue
			}

			// Envía el uso de la CPU por el canal
			if len(percentages) > 0 {
				cpuData <- percentages[0] // Porcentaje de uso total de la CPU
			}
		}
	}
}
