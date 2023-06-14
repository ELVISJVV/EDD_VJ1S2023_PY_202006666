package main

import (
	"EDD_VJ1S2023_PY_202006666/estructuras"
	
	"fmt"
)

func main() {
	
	salirMenu := true
	var opcionMenu string
	var usuario string
	var password string


	
	// Estructuras
	listaEmpleados := &estructuras.ListaSimple{}
	// colaPendientes := Estructuras.Cola{}
	// colaPendientes.ConstructorCola(nil,0)

	// listaDoble := Estructuras.NewLista()

	// pilaAdmin := Estructuras.Pila{}
	// pilaAdmin.ConstructorPila(nil, 0)
	// pilaUsuarios := Estructuras.Pila{}
	// pilaUsuarios.ConstructorPila(nil, 0)


	for salirMenu  {
		fmt.Println("***************** LOGIN ***********************")
		fmt.Println("*            1. Iniciar Sesion                *")
		fmt.Println("*            2. Salir del Sistema             *")
		fmt.Println("***********************************************")
		fmt.Println("")
		fmt.Scanln(&opcionMenu)
		
		if opcionMenu == "1"{
			fmt.Println("")
			fmt.Println("Ingresa tu Usuario")
			fmt.Scanln(&usuario)
			fmt.Println("Ingresa tu Password")
			fmt.Scanln(&password)
			fmt.Println("")

			if usuario =="ADMIN_202006666" && password == "admin"{
				
				salirAdmin := true

				for salirAdmin {
					fmt.Println("***************** Administrador *****************")
					fmt.Println("*            1. Cargar Empleados")
					fmt.Println("*            2. Cargar Imagenes")
					fmt.Println("*            3. Cargar Clientes")
					fmt.Println("*            4. Actualizar Cola")
					fmt.Println("*            5. Reportes Estructuras")
					fmt.Println("*            6. Cerrar Sesion")
					fmt.Println("*************************************************")
					fmt.Println("")
					var opcionAdmin string
					fmt.Scanln(&opcionAdmin)
					if opcionAdmin == "1"{
						fmt.Println("*************** Carga Masiv Empleados **************")
						fmt.Println("Ingrese ruta de la carga masiva")
						var cargaM string
						fmt.Scanln(&cargaM)
						estructuras.LeerArchivoEmpleados(cargaM,listaEmpleados)
						fmt.Println("")
						
					}else if opcionAdmin == "2"{
						fmt.Println("***************** Listado de  Estudiantes *****************")
						// a := listaDoble.SizeLista()
						// b := listaDoble.ReturnEstudianteLista(a)
						// fmt.Println(b.GetNombre())
						// fmt.Println(b.GetPass())
						// Estructuras.Ordenamiento_Burbuja(listaDoble)
						// listaDoble.MostrarConsola()
						// fmt.Println("")
					}else if opcionAdmin == "3"{
						
						
					}else if opcionAdmin == "4"{
							
						// fmt.Println("*************** Carga Masiva **************")
						// fmt.Println("Ingrese ruta de la carga masiva")
						// var cargaM string
						// fmt.Scanln(&cargaM)
						// Estructuras.LeerArchivo(cargaM,&colaPendientes)
						// fmt.Println("")
					}else if opcionAdmin == "5"{
						listaEmpleados.GraficarListaSimple()
					}else if opcionAdmin == "6"{
						salirAdmin = false
					}else{
						fmt.Println("Ingresa una opcion valida")
					}
				}
			} 
			
		} else if opcionMenu == "2" {
			fmt.Println("")
			fmt.Println("Saliendo del Sistema...")
			salirMenu = false
			fmt.Println("")
		}else {
			fmt.Println("")
			fmt.Println("Ingresa una opcion valida")
			fmt.Println("")
		}
	}


// 	fmt.Println("*************** Carga Masiva **************")
// 	fmt.Println("Ingrese ruta de la carga masiva")

// 	ruta := "empleados.csv"
// 	lista := &estructuras.ListaSimple{}
// 	estructuras.LeerArchivoEmpleados(ruta,lista)
// 	lista.GraficarListaSimple()
}