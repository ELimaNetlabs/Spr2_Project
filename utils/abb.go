package utils

import "fmt"

type Nodo struct {
	Nombre   string  // Nombre del proceso
	PID      int     // ID del proceso
	CPUUsage float64 // Consumo de CPU (en porcentaje)
	Izq      *Nodo   // Nodo izquierdo
	Der      *Nodo   // Nodo derecho
}

type ABB struct {
	Raiz *Nodo
}

// Inserta según el uso de CPU
func (a *ABB) Insertar(nuevoNodo *Nodo) {
	a.Raiz = insertarRecursivo(a.Raiz, nuevoNodo)
}

func insertarRecursivo(actual *Nodo, nuevo *Nodo) *Nodo {
	if actual == nil {
		return nuevo
	}

	if nuevo.CPUUsage < actual.CPUUsage {
		actual.Izq = insertarRecursivo(actual.Izq, nuevo)
	} else {
		actual.Der = insertarRecursivo(actual.Der, nuevo)
	}
	return actual
}

// Listar los 5 procesos que más consumen (in-order inverso)
func (a *ABB) ListarTop5() {
	contador := 0
	listarTop5Recursivo(a.Raiz, &contador)
}

func listarTop5Recursivo(nodo *Nodo, contador *int) {
	// Si el nodo es nil, terminamos
	if nodo == nil || *contador >= 5 {
		return
	}

	// Recorremos primero el subárbol derecho (mayores valores de CPU)
	listarTop5Recursivo(nodo.Der, contador)

	// Mostrar el nodo si no hemos mostrado 5 procesos aún
	if *contador < 5 {
		*contador++
		fmt.Printf("Proceso: %s, PID: %d, Uso de CPU: %.2f%%\n", nodo.Nombre, nodo.PID, nodo.CPUUsage)
	}

	// Luego recorremos el subárbol izquierdo
	listarTop5Recursivo(nodo.Izq, contador)
}
