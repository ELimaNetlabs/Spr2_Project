package monitor

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/shirou/gopsutil/v3/process"
)

func VerProceso(pid int) {
	pid32 := int32(pid)
	proc, err := process.NewProcess(pid32)
	if err != nil {
		fmt.Printf("Error al obtener el proceso con PID %d: %v\n", pid, err)
		return
	}

	nombre, _ := proc.Name()
	usuario, _ := proc.Username()
	cpuUsage, _ := proc.CPUPercent()
	memInfo, _ := proc.MemoryInfo()
	createTime, _ := proc.CreateTime()
	uptime := fmt.Sprintf("%d segundos", int64((createTime/1000))-int64(createTime))

	fmt.Printf("Información del proceso con PID %d:\n", pid)
	fmt.Printf("Nombre: %s\n", nombre)
	fmt.Printf("Usuario: %s\n", usuario)
	fmt.Printf("Uso de CPU: %.2f%%\n", cpuUsage)
	if memInfo != nil {
		fmt.Printf("Uso de Memoria: %.2f MB\n", float64(memInfo.RSS)/(1024*1024))
	}
	fmt.Printf("Uptime: %s\n", uptime)
}

func RastrearDetalleProceso(pid int) {
	pid32 := int32(pid)
	proc, err := process.NewProcess(pid32)
	if err != nil {
		fmt.Printf("Error al obtener el proceso con PID %d: %v\n", pid, err)
		return
	}

	fmt.Printf("Iniciando rastreo del proceso con PID %d...\n", pid)

	name, _ := proc.Name()
	exePath, _ := proc.Exe()
	cmdline, _ := proc.Cmdline()
	username, _ := proc.Username()
	cwd, _ := proc.Cwd()

	fmt.Printf("Rastreando detalles del proceso con PID %d:\n", pid)
	fmt.Printf("Nombre del proceso: %s\n", name)
	fmt.Printf("Ruta del ejecutable: %s\n", exePath)
	fmt.Printf("Línea de comandos: %s\n", cmdline)
	fmt.Printf("Usuario que lo inició: %s\n", username)
	fmt.Printf("Directorio de trabajo: %s\n", cwd)
	fmt.Println("Ctrl+c para salir")
}

func DarDeBaja(pid int) {
	cmd := exec.Command("kill", fmt.Sprintf("%d", pid))
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		fmt.Println("Error al intentar matar el proceso:", err)
	} else {
		fmt.Println("Proceso eliminado")
	}
	fmt.Println("Ctrl+c para salir")
}
