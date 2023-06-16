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


func (l *ListaDoble) Insertar(imagen *Imagen) {
	
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
	count := 1
	
	for aux != nil {
		// fmt.Println("************** Imagenes Disponibles **********")
		fmt.Println(count,".  ", aux.Imagen.Nombre)
		// fmt.Println(", Capas: ",aux.Imagen.Capas)
		aux = aux.siguiente
		count++
	}
	fmt.Println("********************************************")
}

	
func (l *ListaDoble) ReturnImagen(posicion int) *Imagen{
	actual := l.Inicio
	i := 1
	for actual != nil{
		if posicion == i{
			return actual.Imagen
		}
		actual = actual.siguiente
		i += 1
		}
	 return nil
	}


func (l *ListaDoble)SizeLista() int {
	// size := l.Longitud 
	return l.Longitud
}


func (c *ListaDoble) GraficarListaDoble() {
	nombre_archivo := "./listaDoble.dot"
	nombre_imagen := "listaDoble.jpg"
	texto := "digraph L{\n"
	texto += "node [shape=box fillcolor=\"#FFEDBB\" style = filled]\n"
	texto += "subgraph cluster_p{\n"
	texto += "label=\"Lista Imagenes \"\n"
	texto += "bgcolor = \"#398D9C\"\n"
	
	
	
	
	size := c.Longitud
	
	for i := 1; i <size +1; i++ {
		

		// texto += "Columna" + strconv.Itoa(i) + "[label = \""+ strconv.Itoa(c.ReturnImagen(i).GetUser())+"\n"+c.ReturnImagen(i).GetNombre() +"\", fillcolor=yellow];\n"

		texto += "Columna" + strconv.Itoa(i) + "[label = \""+ (c.ReturnImagen(i).Nombre)+"\n"+"\", fillcolor=yellow];\n"

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
		texto += "Columna"+ strconv.Itoa(i-1) + "-> Columna" +strconv.Itoa(i)+ ";\n"
	}



	texto += "}\n}\n"



	crearArchivo(nombre_archivo)
	escribirArchivoDot(texto, nombre_archivo)
	ejecutar(nombre_imagen, nombre_archivo)
}

