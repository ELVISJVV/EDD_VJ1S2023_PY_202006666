package estructuras;

import "fmt"

import "strconv"


type ListaSimple struct{
	Primero *NodoSimple
	Longitud int
}



func (l *ListaSimple) estaVacia() bool {
	return l.Longitud == 0
}

func (l *ListaSimple) Insertar(empleado *Empleado) {
	if l.estaVacia() {
		nuevoNodo := &NodoSimple{empleado, nil}
		l.Primero = nuevoNodo
		l.Longitud++
	} else {
		nuevoNodo := &NodoSimple{empleado, nil}
		aux := l.Primero
		for aux.siguiente != nil {
			aux = aux.siguiente
		}
		aux.siguiente = nuevoNodo
		l.Longitud++
	}
}

func (l *ListaSimple) Sacar() {
	if l.estaVacia() {
		fmt.Println("La cola no contiene elementos")
	} else {
		l.Primero = l.Primero.siguiente
		l.Longitud--
	}
}

func (l *ListaSimple) MostrarPrimero() {
	if l.Longitud != 0 {
		fmt.Println("")
		fmt.Println("******** Pendietes ",l.Longitud, " ********" )
		fmt.Println("Estudiante Actual: ",l.Primero.Empleado.Nombre)
	} else {
		fmt.Println("No hay Estudiantes Pendientes")
	}

	}

func (l *ListaSimple) ObtenerEmpleado() *Empleado{
	if l.Longitud != 0 {
		
		return l.Primero.Empleado
	}	
	return nil
} 


func (l *ListaSimple) ReturnEmpleadoListaSimple(posicion int) *Empleado{
	actual := l.Primero
	i := 1
	for actual != nil{
		if posicion == i{
			return actual.Empleado
		}
		actual = actual.siguiente
		i += 1
		}
	 return nil
	}


func (c *ListaSimple) GraficarListaSimple() {
	nombre_archivo := "./listaSimple.dot"
	nombre_imagen := "listaSimple.jpg"
	texto := "digraph L{\n"
	texto += "node [shape=box fillcolor=\"#FFEDBB\" style = filled]\n"
	texto += "subgraph cluster_p{\n"
	texto += "label=\"Lista Empleados \"\n"
	texto += "bgcolor = \"#398D9C\"\n"
	
	
	
	
	size := c.Longitud
	
	for i := 1; i <size +1; i++ {
		

		texto += "Columna" + strconv.Itoa(i) + "[label = \""+ "ID: " + (c.ReturnEmpleadoListaSimple(i).ID)+"\n" + "Nombre: "+c.ReturnEmpleadoListaSimple(i).Nombre +"\", fillcolor=yellow];\n"

	}


	texto += "{rank = same;\n"
	
	for i := 1; i <size +1; i++ {
		if i == size{
			
			texto += "Columna" + strconv.Itoa(i) +"}\n"
		}else{
			texto += "Columna" + strconv.Itoa(i) +";\n"
		}

	}
	
	for i := size; i>0 ; i-- {
		if i-1 == 0{
			break

		}
		texto += "Columna"+ strconv.Itoa(i) + "-> Columna" +strconv.Itoa(i-1)+ ";\n"
	}
	
	texto += "}\n}\n"



	crearArchivo(nombre_archivo)
	escribirArchivoDot(texto, nombre_archivo)
	ejecutar(nombre_imagen, nombre_archivo)
}