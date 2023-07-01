package Lista

type NodoLista struct {
	Empleado  *Empleado
	Siguiente *NodoLista
}

type Empleado struct {
	Id_Empleado string
	Password    string
}
