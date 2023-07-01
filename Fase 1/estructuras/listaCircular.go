package estructuras

import (
	"fmt"
	// "strings"
	"strconv"
)

type ListaCircular struct {
	Inicio   *NodoCircular
	Longitud int
}

func (l *ListaCircular) Insertar(cliente *Cliente) {
	if l.Longitud == 0 {
		l.Inicio = &NodoCircular{Cliente: cliente, siguiente: nil}
		l.Inicio.siguiente = l.Inicio
		l.Longitud++
	} else {
		if l.Longitud == 1 {
			l.Inicio.siguiente = &NodoCircular{Cliente: cliente, siguiente: nil}
			l.Longitud++
		} else {
			aux := l.Inicio
			for i := 0; i < l.Longitud-1; i++ {
				aux = aux.siguiente
			}
			aux.siguiente = &NodoCircular{Cliente: cliente, siguiente: nil}
			l.Longitud++
		}
	}
}

func (l *ListaCircular) Mostrar() {
	aux := l.Inicio
	for i := 0; i < l.Longitud; i++ {
		fmt.Println("Nombre: ", aux.Cliente.Nombre, " ID: ", aux.Cliente.ID)
		aux = aux.siguiente
	}
}


func (l *ListaCircular) ReturnCliente(posicion int) *Cliente {
	actual := l.Inicio
	i := 1
	for actual != nil {
		if posicion == i {
			return actual.Cliente
		}
		actual = actual.siguiente
		i += 1
	}
	return nil
}

func (c *ListaCircular) GraficarListaCircular() {
	nombre_archivo := "./listaCircular.dot"
	nombre_imagen := "listaCircular.jpg"
	texto := "digraph L{\n"
	texto += "node [shape=box fillcolor=\"#FFEDBB\" style = filled]\n"
	texto += "subgraph cluster_p{\n"
	texto += "label=\"Lista Clientes \"\n"
	texto += "bgcolor = \"#398D9C\"\n"

	size := c.Longitud

	for i := 1; i < size+1; i++ {

		texto += "Columna" + strconv.Itoa(i) + "[label = \"" + "ID: " + (c.ReturnCliente(i).ID) + "\n" + "Nombre: " + c.ReturnCliente(i).Nombre + "\", fillcolor=yellow];\n"

	}

	texto += "{rank = same;\n"

	for i := 1; i < size+1; i++ {
		if i == size {

			texto += "Columna" + strconv.Itoa(i) + "}\n"
		} else {
			texto += "Columna" + strconv.Itoa(i) + ";\n"
		}

	}

	for i := size; i > 0; i-- {
		if i-1 == 0 {
			break

		}
		// if i == 1{
		// 	texto += "Columna" + strconv.Itoa(i) + "-> Columna" + strconv.Itoa(size) + ";\n"
		// }
		texto += "Columna" + strconv.Itoa(i) + "-> Columna" + strconv.Itoa(i-1) + ";\n"
	}

	texto += "Columna" + strconv.Itoa(size) + "-> Columna" + strconv.Itoa(1) + ";\n"


	texto += "}\n}\n"

	crearArchivo(nombre_archivo)
	escribirArchivoDot(texto, nombre_archivo)
	ejecutar(nombre_imagen, nombre_archivo)
}
