package utils

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

// Inserta segun el uso de cpu
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
