package monitor

import (
	"Spr2_Project/utils"
	"fmt"
	"time"

	"sync"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/process"
)

func MonitoreoCPU(info chan<- float64, wg *sync.WaitGroup, flag <-chan bool, abb *utils.ABB) {
	defer wg.Done()

	for {
		select {
		case <-flag:
			return
		default:

			pCPU, err := cpu.Percent(1*time.Second, false)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}

			if len(pCPU) > 0 {
				info <- pCPU[0]
			}

			proc, err := process.Processes()
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}

			for _, p := range proc {
				cpu, err := p.CPUPercent()
				if err != nil {
					continue
				}

				name, err := p.Name()
				if err != nil {
					continue
				}

				nodo := &utils.Nodo{
					Nombre: name,
					PID:    int(p.Pid),
					CPU:    cpu,
				}

				abb.Insertar(nodo)
			}

		}
	}
}
