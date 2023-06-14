package estructuras;


// import "fmt"
// import "strconv"


// type Cola struct{
// 	Primero *NodoCola
// 	Longitud int
// }
// func (c *Cola) ConstructorCola(primero *NodoCola, longitud int) {
// 	c.Primero = primero
// 	c.Longitud = longitud
// }

// func (c *Cola) estaVacia() bool {
// 	if c.Longitud == 0 {
// 		return true
// 	} else {
// 		return false
// 	}
// }

// func (c *Cola) Encolar(estudiante *Estudiante) {
// 	if c.estaVacia() {
// 		nuevoNodo := &NodoCola{estudiante, nil}
// 		c.Primero = nuevoNodo
// 		c.Longitud++
// 	} else {
// 		nuevoNodo := &NodoCola{estudiante, nil}
// 		aux := c.Primero
// 		for aux.siguiente != nil {
// 			aux = aux.siguiente
// 		}
// 		aux.siguiente = nuevoNodo
// 		c.Longitud++
// 	}
// }

// func (c *Cola) Descolar() {
// 	if c.estaVacia() {
// 		fmt.Println("La cola no contiene elementos")
// 	} else {
// 		c.Primero = c.Primero.siguiente
// 		c.Longitud--
// 	}
// }

// func (c *Cola) MostrarPrimero() {
// 	if c.Longitud != 0 {
// 		fmt.Println("")
// 		fmt.Println("******** Pendietes ",c.Longitud, " ********" )
// 		fmt.Println("Estudiante Actual: ",c.Primero.estudiante.nombre)
// 	} else {
// 		fmt.Println("No hay Estudiantes Pendientes")
// 	}

// 	}

// func (c *Cola) ObtenerEstudiante() *Estudiante{
// 	if c.Longitud != 0 {
// 		// fmt.Println("******** Pendietes ",c.Longitud, " ********" )
// 		// fmt.Println("Estudiante Actual: ",c.Primero.estudiante.nombre)
		
// 		return c.Primero.estudiante
// 	}	
// 	return nil
// } 


// func (l *Cola) ReturnEstudianteCola(posicion int) *Estudiante{
// 	actual := l.Primero
// 	i := 1
// 	for actual != nil{
// 		if posicion == i{
// 			return actual.estudiante
// 		}
// 		actual = actual.siguiente
// 		i += 1
// 		}
// 	 return nil
// 	}


// func (c *Cola) GraficarCola() {
// 	nombre_archivo := "./cola.dot"
// 	nombre_imagen := "cola.jpg"
// 	texto := "digraph L{\n"
// 	texto += "node [shape=box fillcolor=\"#FFEDBB\" style = filled]\n"
// 	texto += "subgraph cluster_p{\n"
// 	texto += "label=\"Lista Estudiantes Pendientes\"\n"
// 	texto += "bgcolor = \"#398D9C\"\n"
	
	
	
	
// 	size := c.Longitud
	
// 	for i := 1; i <size +1; i++ {
// 		// cont,_ := strconv.Atoi(i) 
// 		// numCarnet ,_= strconv.Atoi(c.ReturnEstudiantePila(i).GetUser())
// 		// name ,_= strconv.Atoi(c.ReturnEstudiantePila(i).GetNombre())


// 		texto += "Columna" + strconv.Itoa(i) + "[label = \""+ strconv.Itoa(c.ReturnEstudianteCola(i).GetUser())+"\n"+c.ReturnEstudianteCola(i).GetNombre() +"\", fillcolor=yellow];\n"

// 	}


// 	texto += "{rank = same;\n"
	
// 	for i := 1; i <size +1; i++ {
// 		if i == size{
			
// 			texto += "Columna" + strconv.Itoa(i) +"}\n"
// 		}else{
// 			texto += "Columna" + strconv.Itoa(i) +";\n"
// 		}

// 	}
	
// 	for i := size; i>0 ; i-- {
// 		if i-1 == 0{
// 			break

// 		}
// 		texto += "Columna"+ strconv.Itoa(i) + "-> Columna" +strconv.Itoa(i-1)+ ";\n"
// 	}
	
// 	texto += "}\n}\n"



// 	crearArchivo(nombre_archivo)
// 	escribirArchivoDot(texto, nombre_archivo)
// 	ejecutar(nombre_imagen, nombre_archivo)
// }