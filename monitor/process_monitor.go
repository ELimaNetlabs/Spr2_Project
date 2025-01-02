package monitor

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/shirou/gopsutil/v3/process"
)

func VerProceso(pid int) {
	pid32 := int32(pid)
	proc, err := process.NewProcess(pid32)
	if err != nil {
		fmt.Printf("Error al obtener el proceso con PID %d: %v\n", pid, err)
		return
	}

	name, _ := proc.Name()
	user, _ := proc.Username()
	cpuP, _ := proc.CPUPercent()
	memInfo, _ := proc.MemoryInfo()
	memP, _ := proc.MemoryPercent()
	pPid, _ := proc.Ppid()
	parent, _ := process.NewProcess(int32(pPid))
	parentName, _ := parent.Name()

	fmt.Printf("Información del proceso con PID %d:\n", pid)
	fmt.Printf("Nombre: %s\n", name)
	fmt.Printf("Usuario: %s\n", user)
	fmt.Printf("CPU: %.2f%% | Memoria: %.2f MB (%.2f%%)\n", cpuP, float64(memInfo.RSS)/(1024*1024), memP)
	fmt.Printf("Proceso padre: %s (PID: %d)\n", parentName, pPid)
}

func RastrearProceso(pid int) {
	proc, err := process.NewProcess(int32(pid))
	if err != nil {
		fmt.Printf("Error al obtener el proceso con PID %d: %v\n", pid, err)
		return
	}

	fmt.Printf("Rastreando detalles del proceso con PID %d...\n", pid)

	name, _ := proc.Name()
	exePath, _ := proc.Exe()
	cmdline, _ := proc.Cmdline()
	username, _ := proc.Username()
	cwd, _ := proc.Cwd()

	createTime, _ := proc.CreateTime()
	uptime := time.Now().Unix() - (createTime / 1000)
	createTimeFormatted := time.Unix(createTime/1000, 0).Format("2006-01-02 15:04:05")

	parentPid, _ := proc.Ppid()
	parent, _ := process.NewProcess(int32(parentPid))
	parentName, _ := parent.Name()

	openFiles, _ := proc.OpenFiles()
	conns, _ := proc.Connections()

	fmt.Printf("Nombre: %s | Ruta: %s | Usuario: %s\n", name, exePath, username)
	fmt.Printf("Cmdline: %s | CWD: %s\n", cmdline, cwd)
	fmt.Printf("Creado: %s | Uptime: %ds\n", createTimeFormatted, uptime)
	fmt.Printf("Proceso padre: %s (PID: %d)\n", parentName, parentPid)

	fmt.Println("Archivos abiertos:")
	for _, f := range openFiles {
		fmt.Printf("- %s\n", f.Path)
	}

	fmt.Println("Conexiones:")
	for _, conn := range conns {
		fmt.Printf("- Local: %s:%d -> Remoto: %s:%d (%s)\n", conn.Laddr.IP, conn.Laddr.Port, conn.Raddr.IP, conn.Raddr.Port, conn.Status)
	}
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
