package estructuras

import (
	"fmt"
	"os"
)

func CrearArchivo() {

	error := os.Remove("Archivo.json")
	if error != nil {
		// fmt.Printf("Error eliminando archivo: %v\n", error)
		fmt.Printf("")
	} else {
		// fmt.Println("Eliminado correctamente")
		fmt.Printf("")
	}
	//Verifica que el archivo existe
	var _, err = os.Stat("Archivo.json")
	//Crea el archivo si no existe
	if os.IsNotExist(err) {
		var file, err = os.Create("Archivo.json")
		if err != nil {
			return
		}
		defer file.Close()
	}
	// fmt.Println("Archivo creado exitosamente", "Archivo.json")
}

func EscribirArchivo(contenido string) {
	var file, err = os.OpenFile("Archivo.json", os.O_RDWR, 0644)
	if err != nil {
		return
	}
	defer file.Close()
	// Escribe algo de texto linea por linea
	_, err = file.WriteString(contenido)
	if err != nil {
		return
	}
	// Salva los cambios
	err = file.Sync()
	if err != nil {
		return
	}
	// fmt.Println("Archivo actualizado existosamente.")
}

func ArchivoJSON(lista *Pila) string {
	contenido := "{\n"
	contenido += "\t\"pedidos\": [\n"
	aux := lista.Primero
	for aux.siguiente != nil {
		contenido += "\t\t{\n"
		contenido += "\t\t\t\"id_cliente\": \"" + aux.Pedido.ID + "\", \n"
		contenido += "\t\t\t\"imagen\": \"" + aux.Pedido.Imagen + "\"\n"
		contenido += "\t\t},\n"
		aux = aux.siguiente
	}
	//esto es para el ultimo elemento
	contenido += "\t\t{\n"
	contenido += "\t\t\t\"id_cliente\": \"" + aux.Pedido.ID + "\", \n"
	contenido += "\t\t\t\"imagen\": \"" + aux.Pedido.Imagen + "\" \n"
	contenido += "\t\t}\n"
	contenido += "\t]\n"
	contenido += "}"
	return contenido
}
