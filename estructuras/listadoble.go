package estructuras


import "fmt"
import "strconv"

type ListaDoble struct {
	Inicio *NodoDoble
	Longitud int
}

func (l *ListaDoble) estaVacia() bool {
	return l.Longitud == 0
}


func (l *ListaDoble) InsertarAlFinal(imagen *Imagen) {
	
	if l.estaVacia() {
		nuevoNodo := &NodoDoble{imagen, nil, nil}
		l.Inicio = (nuevoNodo)
		l.Longitud++
	} else {
		nuevoNodo := &NodoDoble{imagen, nil, nil}
		aux := l.Inicio
		for aux.siguiente != nil {
			aux = aux.siguiente
		}
		
		aux.siguiente = nuevoNodo
		aux.siguiente.anterior = aux
		l.Longitud++
	}
}


func (l *ListaDoble) MostrarConsola() {
	aux := l.Inicio
	for aux != nil {
		fmt.Print("Imagen: ", aux.Imagen.Nombre)
		fmt.Println(", Capas: ",aux.Imagen.Capas)
		fmt.Println("********************************************")
		aux = aux.siguiente
	}
}


func (l *ListaDoble) ReturnEstudianteLista(posicion int) *Estudiante{
	actual := l.Inicio
	i := 1
	for actual != nil{
		if posicion == i{
			return actual.alumno
		}
		actual = actual.siguiente
		i += 1
		}
	 return nil
	}


func (l *ListaDoble)SizeLista() int {
	
	size := l.Longitud 
	return size
}


func (c *ListaDoble) GraficarListaDoble() {
	nombre_archivo := "./lista.dot"
	nombre_imagen := "lista.jpg"
	texto := "digraph L{\n"
	texto += "node [shape=box fillcolor=\"#FFEDBB\" style = filled]\n"
	texto += "subgraph cluster_p{\n"
	texto += "label=\"Lista Estudiantes Pendientes\"\n"
	texto += "bgcolor = \"#398D9C\"\n"
	
	
	
	
	size := c.Longitud
	
	for i := 1; i <size +1; i++ {
		

		texto += "Columna" + strconv.Itoa(i) + "[label = \""+ strconv.Itoa(c.ReturnEstudianteLista(i).GetUser())+"\n"+c.ReturnEstudianteLista(i).GetNombre() +"\", fillcolor=yellow];\n"

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
	


	
	
	for  i:= 1; i< size+1 ; i++  {
		if c.ReturnEstudianteLista(i).GetPila().Longitud != 0{
			texto += "Columna0"+strconv.Itoa(i)+"[shape=record label = \"{|"
			new_size :=c.ReturnEstudianteLista(i).GetPila().Longitud
		for j := 1; j< new_size+1 ; j++  {
		

			texto += c.ReturnEstudianteLista(i).GetPila().ReturnHoraPila(j) + "|"

		}
		texto += "}\", fillcolor=yellow];\n"
		texto += "Columna"+ strconv.Itoa(i) + "-> Columna0" +strconv.Itoa(i)+ ";\n"
		}
		
	}


	
	




	texto += "}\n}\n"



	crearArchivo(nombre_archivo)
	escribirArchivoDot(texto, nombre_archivo)
	ejecutar(nombre_imagen, nombre_archivo)
}

