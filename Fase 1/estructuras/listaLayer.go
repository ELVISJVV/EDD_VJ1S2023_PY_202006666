package estructuras

import "fmt"

// import "strconv"

type ListaLayer struct {
	Primero  *NodoConfig
	Longitud int
}

func (l *ListaLayer) estaVacia() bool {
	return l.Longitud == 0
}

func (l *ListaLayer) Insertar(layer string) {
	if l.estaVacia() {
		nuevoNodo := &NodoConfig{layer, nil}
		l.Primero = nuevoNodo
		l.Longitud++
	} else {
		nuevoNodo := &NodoConfig{layer, nil}
		aux := l.Primero
		for aux.Siguiente != nil {
			aux = aux.Siguiente
		}
		aux.Siguiente = nuevoNodo
		l.Longitud++
	}
}

func (l *ListaLayer) Sacar() {
	if l.estaVacia() {
		fmt.Println("La lsita no contiene elementos")
	} else {
		l.Primero = l.Primero.Siguiente
		l.Longitud--
	}
}

func (l *ListaLayer) MostrarLayer() {
	aux := l.Primero
	count := 1

	fmt.Println("************** Capas Disponibles **********")
	for aux != nil {
		fmt.Println(count, ".  ", aux.Layer)
		// fmt.Println(", Capas: ",aux.Imagen.Capas)
		aux = aux.Siguiente
		count++
	}
	fmt.Println("********************************************")
}

func (l *ListaLayer) ObtenerLayer() string {
	if l.Longitud != 0 {

		return l.Primero.Layer
	}
	return ""
}

func (l *ListaLayer) ReturnLayer(posicion int) string {
	actual := l.Primero
	i := 1
	for actual != nil {
		if posicion == i {
			return actual.Layer
		}
		actual = actual.Siguiente
		i += 1
	}
	return ""
}
