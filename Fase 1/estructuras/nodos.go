package estructuras

type NodoCola struct {
	Cliente    *Cliente
	siguiente  *NodoCola
}


type NodoPila struct {
	Pedido     *Pedido
	siguiente  *NodoPila
}

type NodoCircular struct {
	Cliente   *Cliente
	siguiente *NodoCircular
}

type NodoSimple struct {
	Empleado   *Empleado
	siguiente  *NodoSimple
}

type NodoDoble struct {
	Imagen    *Imagen
	siguiente *NodoDoble
	anterior  *NodoDoble
}

type NodoMatriz struct {
	Siguiente *NodoMatriz
	Anterior  *NodoMatriz
	Abajo     *NodoMatriz
	Arriba    *NodoMatriz
	PosX      int
	PosY      int
	Color     string
}

type NodoConfig struct{
	Layer string
	Siguiente *NodoConfig
}

