package estructuras

import (
	"math/rand"
	"strconv"
)

func Verificar(colaActual *Cola, clientes *ListaCircular) bool {
	aux := colaActual.Primero
	if aux != nil {
		if aux.Cliente.ID != "X" {
			for i := 1; i < clientes.Longitud+1; i++ {
				if aux.Cliente.ID == clientes.ReturnCliente(i).ID {
					return true
				}
			}
		} else {
			return false
		}
	}
	return false
}

func AsignarLista(colaActual *Cola, clientes *ListaCircular) string {
	valorString := ValidarRepetido(clientes)

	cliente := Cliente{Nombre: colaActual.Primero.Cliente.Nombre, ID: valorString}

	clientes.Insertar(&cliente)

	cadena := "Se asigno el ID: " + valorString + " al cliente " + colaActual.Primero.Cliente.Nombre
	colaActual.Primero.Cliente.ID = valorString
	
	return cadena
}

func ValidarRepetido(clientes *ListaCircular) string {
	valor := (rand.Intn(8000)) + 1000 // + 1000
	valorString := strconv.Itoa(valor)
	validacion := true
	for validacion {
		validacion = false
		valor = (rand.Intn(8000)) + 1000
		valorString = strconv.Itoa(valor)
		for i := 1; i < clientes.Longitud+1; i++ {
			if clientes.ReturnCliente(i).ID == valorString {
				validacion = true
			}
		}
	}

	return valorString

}
