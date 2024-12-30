package utils

import "fmt"

type Nodo struct {
	Nombre   string
	PID      int
	CPUUsage float64
	Izq      *Nodo
	Der      *Nodo
}

type ABB struct {
	Raiz *Nodo
}

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

func (a *ABB) ListarTop5() {
	contador := 0
	listarTop5Recursivo(a.Raiz, &contador)
}

func listarTop5Recursivo(nodo *Nodo, contador *int) {

	if nodo == nil || *contador >= 5 {
		return
	}

	listarTop5Recursivo(nodo.Der, contador)

	if *contador < 5 {
		*contador++
		fmt.Printf("Proceso: %s, PID: %d, Uso de CPU: %.2f%%\n", nodo.Nombre, nodo.PID, nodo.CPUUsage)
	}

	listarTop5Recursivo(nodo.Izq, contador)
}
