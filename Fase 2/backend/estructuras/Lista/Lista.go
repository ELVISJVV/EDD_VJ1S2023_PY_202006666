package Lista

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

type ListaSimple struct {
	Inicio   *NodoLista
	Longitud int
}

func (l *ListaSimple) estaVacia() bool {
	return l.Longitud == 0
}

func (l *ListaSimple) Insertar(carnet string, password string) {
	nuevoEmpleado := &Empleado{Id_Empleado: carnet, Password: password}
	if l.estaVacia() {
		l.Inicio = &NodoLista{Empleado: nuevoEmpleado}
		l.Longitud++
	} else {
		aux := l.Inicio
		for aux.Siguiente != nil {
			aux = aux.Siguiente
		}
		aux.Siguiente = &NodoLista{Empleado: nuevoEmpleado}
		l.Longitud++
	}
}

func (l *ListaSimple) LeerArchivo(ruta string) {
	file, err := os.Open(ruta)
	if err != nil {
		fmt.Println("No pude abrir el archivo")
		return
	}
	defer file.Close()

	lectura := csv.NewReader(file)
	lectura.Comma = ','
	encabezado := true
	for {
		linea, err := lectura.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("No pude leer la linea del csv")
			continue
		}
		if encabezado {
			encabezado = false
			continue
		}
		l.Insertar(linea[0], linea[3])
	}
}

func (l *ListaSimple) Buscar(id_empleado string, password string) bool {
	aux := l.Inicio
	for aux != nil {
		if aux.Empleado.Id_Empleado == id_empleado && aux.Empleado.Password == password {
			return true
		}
		aux = aux.Siguiente
	}
	return false
}
