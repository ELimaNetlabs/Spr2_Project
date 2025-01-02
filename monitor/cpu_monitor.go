package monitor

import (
	"Spr2_Project/utils"
	"fmt"
	"time"

	"sync"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/process"
)

func MonitoreoCPU(data chan<- float64, wg *sync.WaitGroup, done <-chan bool, abb *utils.ABB) {
	defer wg.Done()

	for {
		select {
		case <-done:
			return
		default:

			percentages, err := cpu.Percent(1*time.Second, false)
			if err != nil {
				fmt.Println("Error obteniendo uso de CPU:", err)
				continue
			}

			if len(percentages) > 0 {
				data <- percentages[0]
			}

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

				name, err := p.Name()
				if err != nil {
					continue
				}

				nodo := &utils.Nodo{
					Nombre:   name,
					PID:      int(p.Pid),
					CPUUsage: cpuUsage,
				}

				abb.Insertar(nodo)
			}

		}
	}
}
