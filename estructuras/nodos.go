package estructuras

// type NodoCola struct {
// 	Cliente    *Cliente
// 	siguiente  *NodoCola
// }


// type NodoPila struct {
// 	Pedido     *Pedido
// 	siguiente  *NodoPila
// }

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
