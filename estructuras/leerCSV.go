package estructuras

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func LeerArchivoEmpleados(ruta string, listaAux *ListaSimple) {
	//listaAux := &ListaCircular{Inicio: nil, Longitud: 0}
	file, err := os.Open(ruta)
	if err != nil {
		fmt.Println("No pude abrir el archivo")
		return
	}
	defer file.Close()

	lectura := csv.NewReader(file)
	lectura.Comma = ','
	encabezado := true
	for {
		linea, err := lectura.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("No pude leer la linea del csv")
			continue
		}
		if encabezado {
			encabezado = false
			continue
		}
		
		empleado := Empleado{Nombre: linea[1], ID: linea[0], Cargo: linea[2], Contraseña: linea[3]}	

		listaAux.Insertar(&empleado)
	}
}



func LeerArchivoImagenes(ruta string, listaAux *ListaDoble) {
	
	file, err := os.Open(ruta)
	if err != nil {
		fmt.Println("No pude abrir el archivo")
		return
	}
	defer file.Close()

	lectura := csv.NewReader(file)
	lectura.Comma = ','
	encabezado := true
	for {
		linea, err := lectura.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("No pude leer la linea del csv")
			continue
		}
		if encabezado {
			encabezado = false
			continue
		}
		
		imagen := Imagen{Nombre: linea[0], Capas: linea[1]}	

		listaAux.Insertar(&imagen)
	}
}


func LeerArchivoClientes(ruta string, listaAux *ListaCircular) {
	
	file, err := os.Open(ruta)
	if err != nil {
		fmt.Println("No pude abrir el archivo")
		return
	}
	defer file.Close()

	lectura := csv.NewReader(file)
	lectura.Comma = ','
	encabezado := true
	for {
		linea, err := lectura.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("No pude leer la linea del csv")
			continue
		}
		if encabezado {
			encabezado = false
			continue
		}
		
		cliente := Cliente{Nombre: linea[1], ID: linea[0]}	

		listaAux.Insertar(&cliente)
	}
}
