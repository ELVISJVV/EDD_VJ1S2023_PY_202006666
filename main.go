package main

import (
	"EDD_VJ1S2023_PY_202006666/estructuras"
	"fmt"
	"os/exec"
	"strconv"
)

func main() {

	salirMenu := true
	var opcionMenu string
	var usuario string
	var password string

	// Estructuras
	listaEmpleados := &estructuras.ListaSimple{}
	listaImagenes := &estructuras.ListaDoble{}
	listaClientes := &estructuras.ListaCircular{}
	colaClientes := &estructuras.Cola{}
	pilaPedidos := &estructuras.Pila{}
	matriz := &estructuras.Matriz{Raiz: &estructuras.NodoMatriz{PosX: -1, PosY: -1, Color: "RAIZ"}}

	for salirMenu {
		fmt.Println("***************** LOGIN ***********************")
		fmt.Println("*            1. Iniciar Sesion                *")
		fmt.Println("*            2. Salir del Sistema             *")
		fmt.Println("***********************************************")
		fmt.Println("")
		fmt.Scanln(&opcionMenu)

		if opcionMenu == "1" {
			fmt.Println("")
			fmt.Println("Ingresa tu Usuario")
			fmt.Scanln(&usuario)
			fmt.Println("Ingresa tu Password")
			fmt.Scanln(&password)
			fmt.Println("")

			if usuario == "ADMIN_202006666" && password == "admin" {

				salirAdmin := true

				for salirAdmin {
					fmt.Println("************* Administrador 202006666 *************")
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
					if opcionAdmin == "1" {
						fmt.Println("*************** Cargar Empleados **************")
						fmt.Println("Ingrese ruta de la carga de empleados")
						var cargaM string
						fmt.Scanln(&cargaM)
						estructuras.LeerArchivoEmpleados(cargaM, listaEmpleados)
						fmt.Println("")

					} else if opcionAdmin == "2" {
						fmt.Println("*************** Cargar Imagenes **************")
						fmt.Println("Ingrese ruta de la carga de imagenes")
						var cargaM string
						fmt.Scanln(&cargaM)
						estructuras.LeerArchivoImagenes(cargaM, listaImagenes)
						fmt.Println("")
					} else if opcionAdmin == "3" {
						fmt.Println("*************** Cargar Clientes **************")
						fmt.Println("Ingrese ruta de la carga de clientes")
						var cargaM string
						fmt.Scanln(&cargaM)
						estructuras.LeerArchivoClientes(cargaM, listaClientes)
						fmt.Println("")
					} else if opcionAdmin == "4" {

						fmt.Println("*************** Actualizar  Cola **************")
						fmt.Println("Ingrese ruta de la carga de la cola")
						var cargaM string
						fmt.Scanln(&cargaM)
						estructuras.LeerArchivoActualizarCola(cargaM, colaClientes)
						// Estructuras.LeerArchivo(cargaM,&colaPendientes)
						fmt.Println("")
					} else if opcionAdmin == "5" {
						salirReportes := true
						for salirReportes {
							fmt.Println("***************** Reportes *****************")
							fmt.Println("*            1. Reporte Empleados")
							fmt.Println("*            2. Reporte Imagenes")
							fmt.Println("*            3. Reporte Clientes")
							fmt.Println("*            4. Reporte Cola")
							fmt.Println("*            5. Reporte Pila")
							fmt.Println("*            6. Reporte Capas")
							fmt.Println("*            7. Regresar")
							fmt.Println("*********************************************")
							fmt.Println("")
							var opcionReportes string
							fmt.Scanln(&opcionReportes)
							if opcionReportes == "1" {
								listaEmpleados.GraficarListaSimple()
								cmd := exec.Command("cmd", "/c", "start", "listaSimple.jpg")
								err := cmd.Start()
								if err != nil {
									fmt.Println("No pude abrir imagen")
								}
								fmt.Println("")
							} else if opcionReportes == "2" {
								listaImagenes.GraficarListaDoble()
								cmd := exec.Command("cmd", "/c", "start", "listaDoble.jpg")
								err := cmd.Start()
								if err != nil {
									fmt.Println("No pude abrir imagen")
								}
								fmt.Println("")
							} else if opcionReportes == "3" {
								listaClientes.GraficarListaCircular()
								cmd := exec.Command("cmd", "/c", "start", "listaCircular.jpg")
								err := cmd.Start()
								if err != nil {
									fmt.Println("No pude abrir imagen")
								}
								fmt.Println("")
							} else if opcionReportes == "4" {
								colaClientes.GraficarCola()
								cmd := exec.Command("cmd", "/c", "start", "cola.jpg")
								err := cmd.Start()
								if err != nil {
									fmt.Println("No pude abrir imagen")
								}
								fmt.Println("")
							} else if opcionReportes == "5" {
								pilaPedidos.GraficarPila()
								cmd := exec.Command("cmd", "/c", "start", "pilaPedidos.jpg")
								err := cmd.Start()
								if err != nil {
									fmt.Println("No pude abrir imagen")
								}
								fmt.Println("")
							} else if opcionReportes == "6" {
								salirVerImagenes := true
								for salirVerImagenes {
									fmt.Println("********** Ver  Imagenes Cargadas *********")
									if listaImagenes.Longitud == 0 {
										fmt.Println("No hay imagenes cargadas")
										fmt.Println("")
										salirVerImagenes = false
										break
									}
									listaImagenes.MostrarConsola()
									fmt.Println("")
									fmt.Println("Eliga una imagen de la lista:")
									var imagen string
									fmt.Scanln(&imagen)
									fmt.Println("")

									i, _ := strconv.Atoi(imagen)
									if i >= 1 && i <= listaImagenes.Longitud {
										salirVerImagenes = false
										imagenElegida := listaImagenes.ReturnImagen(i).Nombre
										listaCapas := &estructuras.ListaLayer{}
										estructuras.LeerArchivoConfig("csv/"+imagenElegida+"/inicial.csv", listaCapas)
										salircapas := true
										for salircapas {
											listaCapas.MostrarLayer()
											fmt.Println("Eliga una capa de la lista:")
											var capa string
											fmt.Scanln(&capa)
											fmt.Println("")
											j, _ := strconv.Atoi(capa)
											if j >= 1 && j <= listaCapas.Longitud {
												capaElegida := listaCapas.ReturnLayer(j)

												matriz = &estructuras.Matriz{Raiz: &estructuras.NodoMatriz{PosX: -1, PosY: -1, Color: "RAIZ"}}
												matriz.LeerArchivo("csv/" + imagenElegida + "/" + capaElegida + ".csv")
												matriz.Reporte()
												cmd := exec.Command("cmd", "/c", "start", "matriz.jpg")
								err := cmd.Start()
								if err != nil {
									fmt.Println("No pude abrir imagen")
								}
												salircapas = false
											} else {
												fmt.Println("Ingrese una opcion valida")
											}
										}

									} else {
										fmt.Println("Ingrese una opcion valida")
									}
								}
							} else if opcionReportes == "7" {
								salirReportes = false
							} else {
								fmt.Println("Ingresa una opcion valida")
							}
						}

					} else if opcionAdmin == "6" {
						salirAdmin = false
					} else {
						fmt.Println("Ingresa una opcion valida")
					}
				}
			} else {
				if listaEmpleados.Longitud == 0 {
					fmt.Println("No se encontraron coincidencias")
				} else {
					validacionEmpleados := false
					size := listaEmpleados.Longitud + 1
					for i := 1; i < size; i++ {
						if listaEmpleados.ReturnEmpleadoListaSimple(i).ID == usuario && listaEmpleados.ReturnEmpleadoListaSimple(i).Contraseña == password {
							fmt.Println("Se inicio correctamente")
							validacionEmpleados = true
							salirEmpleado := true
							for salirEmpleado {
								fmt.Println("")
								fmt.Println("*************** Empleado " + listaEmpleados.ReturnEmpleadoListaSimple(i).ID + " ***************")
								fmt.Println("*            1. Ver imagenes Cargadas")
								fmt.Println("*            2. Realizar  Pedido")
								fmt.Println("*            3. Cerrar Sesion")
								fmt.Println("*********************************************")
								fmt.Println("")
								var opcionEmpleado string
								fmt.Scanln(&opcionEmpleado)
								if opcionEmpleado == "1" {
									salirVerImagenes := true
									for salirVerImagenes {
										fmt.Println("********** Ver  Imagenes Cargadas *********")
										if listaImagenes.Longitud == 0 {
											fmt.Println("No hay imagenes cargadas")
											fmt.Println("")
											salirVerImagenes = false
											break
										}
										listaImagenes.MostrarConsola()
										fmt.Println("")
										fmt.Println("Eliga una imagen de la lista:")
										// listaImagenes.MostrarConsola()
										var imagen string
										fmt.Scanln(&imagen)
										i, _ := strconv.Atoi(imagen)
										if i >= 1 && i <= listaImagenes.Longitud {
											salirVerImagenes = false
											imagenElegida := listaImagenes.ReturnImagen(i).Nombre
											matriz_csv := &estructuras.Matriz{Raiz: &estructuras.NodoMatriz{PosX: -1, PosY: -1, Color: "RAIZ"}}
											matriz_csv.LeerInicial("csv/"+imagenElegida+"/inicial.csv", imagenElegida)
											matriz_csv.GenerarImagen(imagenElegida)
											cmd := exec.Command("cmd", "/c", "start", "csv/" + imagenElegida + "/" + imagenElegida + ".html")
								err := cmd.Start()
								if err != nil {
									fmt.Println("No pude abrir imagen")
								}

										} else {
											fmt.Println("Ingrese una opcion valida")
										}
									}

								} else if opcionEmpleado == "2" {
									fmt.Println("*************** Realizar Pedido **************")

									existe := estructuras.Verificar(colaClientes, listaClientes)
									if existe && colaClientes.Primero != nil { /*Usuario ya esta registrado en lista circular y la Cola aun tiene elementos*/

										validar := true
										for validar {
											fmt.Println("Cliente atendido por el empleado " + listaEmpleados.ReturnEmpleadoListaSimple(i).ID)
											fmt.Println("El ID del cliente actual es: ", colaClientes.Primero.Cliente.ID)
											fmt.Println("Eliga una imagen de la lista:")
											if listaImagenes.Longitud == 0 {
												fmt.Println("No hay imagenes cargadas")
												fmt.Println("")
												validar = false
												break
											}
											listaImagenes.MostrarConsola()
											var imagen string
											fmt.Scanln(&imagen)
											i, _ := strconv.Atoi(imagen)
											if i >= 1 && i <= listaImagenes.Longitud {
												validar = false
												imagenElegida := listaImagenes.ReturnImagen(i).Nombre
												pedido := estructuras.Pedido{ID: colaClientes.Primero.Cliente.ID, Imagen: imagenElegida}
												pilaPedidos.Push(&pedido)
												colaClientes.Descolar()

											} else {
												fmt.Println("Ingrese una opcion valida")
											}
										}
										if pilaPedidos.Primero != nil {
											contenido := estructuras.ArchivoJSON(pilaPedidos)
											estructuras.CrearArchivo()
											estructuras.EscribirArchivo(contenido)
										}

									} else if !existe && colaClientes.Primero != nil {
										fmt.Println("Cliente atendido por el empleado " + listaEmpleados.ReturnEmpleadoListaSimple(i).ID)
										variable := estructuras.AsignarLista(colaClientes, listaClientes)
										fmt.Println(variable)
										validar := true
										for validar {

											fmt.Println("Eliga una imagen de la lista:")
											if listaImagenes.Longitud == 0 {

												fmt.Println("")
												fmt.Println("No hay imagenes cargadas")
												fmt.Println("")
												validar = false
												break
											}
											listaImagenes.MostrarConsola()
											var imagen string
											fmt.Scanln(&imagen)
											i, _ := strconv.Atoi(imagen)
											if i >= 1 && i <= listaImagenes.Longitud {
												validar = false
												imagenElegida := listaImagenes.ReturnImagen(i).Nombre
												pedido := estructuras.Pedido{ID: colaClientes.Primero.Cliente.ID, Imagen: imagenElegida}
												pilaPedidos.Push(&pedido)
												colaClientes.Descolar()

											}
										}
										if pilaPedidos.Primero != nil {

											contenido := estructuras.ArchivoJSON(pilaPedidos)
											estructuras.CrearArchivo()
											estructuras.EscribirArchivo(contenido)
										}
									} else if colaClientes.Primero == nil {
										fmt.Println("Ya no hay alumnos por atender")
									}

								} else if opcionEmpleado == "3" {
									salirEmpleado = false

								} else {
									fmt.Println("Ingresa una opcion valida")
								}
							}

							break
						}

					}
					if !validacionEmpleados {
						fmt.Println("No se encontraron coincidencias")
					}
				}
			}

		} else if opcionMenu == "2" {
			fmt.Println("")
			fmt.Println("Saliendo del Sistema...")
			salirMenu = false
			fmt.Println("")
		} else {
			fmt.Println("")
			fmt.Println("Ingresa una opcion valida")
			fmt.Println("")
		}
	}

}
