package monitor

import (
	"fmt"
	"sync"
	"time"

	"github.com/shirou/gopsutil/v3/process"
)

func VerProceso(info chan<- string, wg2 *sync.WaitGroup, pid int) {
	defer wg2.Done()

	pid2 := int32(pid)
	proc, err := process.NewProcess(pid2)
	if err != nil {
		info <- fmt.Sprintf("Error al obtener el proceso con PID %d: %v", pid, err)
		return
	}

	// Recopilar información básica
	nombre, _ := proc.Name()
	usuario, _ := proc.Username()
	cpuUsage, _ := proc.CPUPercent()
	memInfo, _ := proc.MemoryInfo()
	createTime, _ := proc.CreateTime()
	uptime := fmt.Sprintf("%d segundos", (int64((createTime / 1000)) - int64(createTime)))

	// Formatear la información en un string
	data := fmt.Sprintf(
		"Información del proceso con PID %d:\nNombre: %s\nUsuario: %s\nUso de CPU: %.2f%%\n",
		pid, nombre, usuario, cpuUsage,
	)
	if memInfo != nil {
		data += fmt.Sprintf("Uso de Memoria: %.2f MB\n", float64(memInfo.RSS)/(1024*1024))
	}
	data += fmt.Sprintf("Uptime: %s\n", uptime)

	// Enviar la información al canal
	info <- data
}

// RastrearDetalleProceso obtiene detalles avanzados de un proceso.
func RastrearDetalleProceso(pid int, done chan bool) {
	pid32 := int32(pid)
	proc, err := process.NewProcess(pid32)
	if err != nil {
		fmt.Printf("Error al obtener el proceso con PID %d: %v\n", pid, err)
		return
	}

	fmt.Printf("Iniciando rastreo del proceso con PID %d...\n", pid)
	fmt.Println("Presiona Ctrl+C para detener el rastreo.")

	// Bucle de monitoreo
	for {
		select {
		case <-done:
			fmt.Println("Rastreo detenido.")
			return
		default:
			// Obtener información del proceso
			name, _ := proc.Name()       // Nombre del proceso
			exePath, _ := proc.Exe()     // Ruta completa del ejecutable
			cmdline, _ := proc.Cmdline() // Línea de comandos usada para iniciar el proceso
			//parent, _ := proc.Parent()     // Proceso padre
			username, _ := proc.Username() // Usuario que inició el proceso
			cwd, _ := proc.Cwd()           // Directorio de trabajo actual del proceso

			fmt.Print("\033[H\033[2J") // Limpiar pantalla
			fmt.Printf("Rastreando detalles del proceso con PID %d:\n", pid)
			fmt.Printf("Nombre del proceso: %s\n", name)
			fmt.Printf("Ruta del ejecutable: %s\n", exePath)
			fmt.Printf("Línea de comandos: %s\n", cmdline)
			// if parent != nil {
			// 	fmt.Printf("Proceso padre: %s (PID: %d)\n", parent.Name, parent.Pid)
			// }
			fmt.Printf("Usuario que lo inició: %s\n", username)
			fmt.Printf("Directorio de trabajo: %s\n", cwd)

			time.Sleep(2 * time.Second) // Intervalo de actualización
		}
	}
}
