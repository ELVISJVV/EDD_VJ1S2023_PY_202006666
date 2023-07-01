package estructuras


type Empleado struct {
	Nombre     string
	ID   	   string
	Cargo      string
	Contraseña string
	
}

func (e *Empleado) ConstructorEmpleado(nombre string, id string, password string, cargo string) {
	e.Nombre = nombre
	e.ID = id
	e.Contraseña = password
	e.Cargo = cargo
}



