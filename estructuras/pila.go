package estructuras

import (
	"fmt"
)

type Pila struct {
	Primero  *NodoPila
	Longitud int
}

func (c *Pila) ConstructorPila(primero *NodoPila, longitud int) {
	c.Primero = primero
	c.Longitud = longitud
}

func (p *Pila) estaVacia() bool {
	if p.Longitud == 0 {
		return true
	} else {
		return false
	}
}

func (p *Pila) Push(pedido *Pedido) {
	if p.estaVacia() {
		nuevoNodo := &NodoPila{pedido, nil}
		p.Primero = nuevoNodo
		p.Longitud++
	} else {
		nuevoNodo := &NodoPila{pedido, p.Primero}
		p.Primero = nuevoNodo
		p.Longitud++
	}
}

func (p *Pila) Pop() {
	if p.estaVacia() {
		fmt.Println("La pila no tiene elementos")
	} else {
		p.Primero = p.Primero.siguiente
		p.Longitud--
	}
}

func (p *Pila) Peek() {
	if p.estaVacia() {
		fmt.Println("La pila no tiene elementos")
	} else {

		fmt.Println(p.Primero.Pedido.ID)
	}
}


func (p *Pila) MostrarPila() {
	aux := p.Primero
	for aux != nil {
		fmt.Println(aux.Pedido)
		aux = aux.siguiente
	}
}

func (l *Pila) ReturnPedido(posicion int) *Pedido{
	actual := l.Primero
	i := 1
	for actual != nil{
		if posicion == i{
			return actual.Pedido
		}
		actual = actual.siguiente
		i += 1
		}
	return nil
	}







func (c *Pila) GraficarPila() {
	nombre_archivo := "./pilaPedidos.dot"
	nombre_imagen := "pilaPedidos.jpg"
	texto := "digraph L{\n"
	texto += "node [shape=box fillcolor=\"#FFEDBB\" style = filled]\n"
	texto += "subgraph cluster_p{\n"
	texto += "label=\"Pila Administrador\"\n"
	texto += "bgcolor = \"#398D9C\"\n"
	
	
	
	
	size := c.Longitud
	texto += "Columna[shape=record label = \"{|"
	
	
	for i := 1; i< size+1 ; i++  {
		

		texto += c.ReturnPedido(i).ID+"|"

	}

	
	texto += "}\", fillcolor=yellow];\n}\n}"



	crearArchivo(nombre_archivo)
	escribirArchivoDot(texto, nombre_archivo)
	ejecutar(nombre_imagen, nombre_archivo)
}