package main

import (
	"backend/estructuras/ArbolAVL"
	"backend/estructuras/ColaPedidos"
	"backend/estructuras/Facturas"
	"backend/estructuras/Grafo"
	"backend/estructuras/Lista"
	"backend/estructuras/Matriz"
	"backend/estructuras/Peticiones"
	"backend/estructuras/TablaHash"
	"strconv"

	// "backend/estructuras/GenerarArchivos"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os/exec"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var ListaEmpleados *Lista.ListaSimple
var ArbolPedidos *ArbolAVL.Arbol
var MatrizOriginal *Matriz.Matriz
var MatrizFiltro *Matriz.Matriz
var PedidosCola *ColaPedidos.Cola
var FacturasRealizadas *Facturas.BlockChain
var VerFacturasRealizadas *TablaHash.TablaHash
var FiltrosColocados string
var EmpleadoLogeado string
var GrafosEmpleados map[string]Grafo.Grafo
var GrafoEmpleado *Grafo.Grafo

func main() {
	// estructuras utilizadas
	ListaEmpleados = &Lista.ListaSimple{Inicio: nil, Longitud: 0}
	ArbolPedidos = &ArbolAVL.Arbol{Raiz: nil}
	MatrizOriginal = &Matriz.Matriz{Raiz: &Matriz.NodoMatriz{PosX: -1, PosY: -1, Color: "Raiz"}}
	MatrizFiltro = &Matriz.Matriz{Raiz: &Matriz.NodoMatriz{PosX: -1, PosY: -1, Color: "Raiz"}}
	PedidosCola = &ColaPedidos.Cola{Primero: nil, Longitud: 0}
	FacturasRealizadas = &Facturas.BlockChain{Bloques_Creados: 0}
	VerFacturasRealizadas = &TablaHash.TablaHash{Capacidad: 5, Utilizacion: 0}
	FiltrosColocados = ""
	EmpleadoLogeado = ""
	GrafoEmpleado = &Grafo.Grafo{Principal: nil}

	app := fiber.New()
	app.Use(cors.New())

	app.Post("/login", func(c *fiber.Ctx) error {
		var usuario Peticiones.Login
		c.BodyParser(&usuario)
		// fmt.Println(usuario)

		if usuario.Username == "ADMIN_202006666" && usuario.Password == "admin" {
			return c.JSON(&fiber.Map{
				"status": "admin",
			})
		} else {
			if ListaEmpleados.Inicio != nil {
				if ListaEmpleados.Buscar(usuario.Username, usuario.Password) {
					VerFacturasRealizadas = &TablaHash.TablaHash{Capacidad: 5, Utilizacion: 0}
					VerFacturasRealizadas.NewTablaHash()
					EmpleadoLogeado = usuario.Username
					return c.JSON(&fiber.Map{
						"status": "employee",
					})
				}
			} else {
				return c.JSON(&fiber.Map{
					"status": "unknown",
				})
			}
		}

		return c.JSON(&fiber.Map{
			"status": "unknown",
		})
	})

	app.Post("/cargarempleados", func(c *fiber.Ctx) error {
		var nombreArchivo Peticiones.Archivo
		c.BodyParser(&nombreArchivo)
		// fmt.Println(nombreArchivo)
		fmt.Println("")
		ListaEmpleados.LeerArchivo(nombreArchivo.Nombre)
		// fmt.Println(ListaEmpleados.Inicio.Empleado.Id_Cliente)
		return c.JSON(&fiber.Map{
			"status": 200,
		})
	})

	app.Post("/cargarpedidos", func(c *fiber.Ctx) error {
		var pedidosNuevos Peticiones.ArbolPeticion
		c.BodyParser(&pedidosNuevos)
		for i := 0; i < len(pedidosNuevos.Pedidos); i++ {
			ArbolPedidos.InsertarElemento(pedidosNuevos.Pedidos[i].Id_Cliente, pedidosNuevos.Pedidos[i].Nombre_Imagen)
		}
		ArbolPedidos.RecorridoInorden(ArbolPedidos.Raiz, PedidosCola)
		return c.JSON(&fiber.Map{
			"status": 200,
			"cola":   PedidosCola,
		})
	})

	app.Get("/reporte-arbol", func(c *fiber.Ctx) error {
		ArbolPedidos.Graficar()
		var imagen Peticiones.RespuestaImagen = Peticiones.RespuestaImagen{Nombre: "Reporte/arbolAVL.jpg"}
		/*INICIO*/
		imageBytes, err := ioutil.ReadFile(imagen.Nombre)
		if err != nil {
			return c.JSON(&fiber.Map{
				"status": 404,
			})
		}
		// Codifica los bytes de la imagen en base64
		imagen.Imagenbase64 = "data:image/jpg;base64," + base64.StdEncoding.EncodeToString(imageBytes)
		return c.JSON(&fiber.Map{
			"status": 200,
			"imagen": imagen,
		})
	})

	app.Get("/reporte-grafo", func(c *fiber.Ctx) error {
		GrafoEmpleado.Reporte()
		var imagen Peticiones.RespuestaImagen = Peticiones.RespuestaImagen{Nombre: "Reporte/grafo.jpg"}
		/*INICIO*/
		imageBytes, err := ioutil.ReadFile(imagen.Nombre)
		if err != nil {
			return c.JSON(&fiber.Map{
				"status": 404,
			})
		}
		// Codifica los bytes de la imagen en base64
		imagen.Imagenbase64 = "data:image/jpg;base64," + base64.StdEncoding.EncodeToString(imageBytes)
		return c.JSON(&fiber.Map{
			"status": 200,
			"imagen": imagen,
		})
	})

	app.Get("/reporte-bloque", func(c *fiber.Ctx) error {
		FacturasRealizadas.GenerateGraph()
		var imagen Peticiones.RespuestaImagen = Peticiones.RespuestaImagen{Nombre: "blockchain-graph.png"}
		/*INICIO*/
		imageBytes, err := ioutil.ReadFile(imagen.Nombre)
		if err != nil {
			return c.JSON(&fiber.Map{
				"status": 404,
			})
		}
		// Codifica los bytes de la imagen en base64
		imagen.Imagenbase64 = "data:image/jpg;base64," + base64.StdEncoding.EncodeToString(imageBytes)
		return c.JSON(&fiber.Map{
			"status": 200,
			"imagen": imagen,
		})
	})

	app.Post("/aplicarfiltro", func(c *fiber.Ctx) error {
		var tipo Peticiones.PeticionFiltro
		c.BodyParser(&tipo)
		// fmt.Println(tipo)
		if PedidosCola.Primero == nil {
			return c.JSON(&fiber.Map{
				"status": 400,
			})
		}

		tipo.NombreImagen = PedidosCola.Primero.Pedido.Nombre_Imagen
		var nombreFiltro = ""
		var nombreHtml =""
		switch tipo.Tipo {
		case 1:
			nombreFiltro = "Negativo"
			MatrizFiltro.LeerInicial("csv/"+tipo.NombreImagen+"/inicial.csv", tipo.NombreImagen)
			MatrizFiltro.Negativo(tipo.NombreImagen + "Negativo")
			nombreHtml = tipo.NombreImagen + "Negativo"
			FiltrosColocados += "Negativo, "
		case 2:
			nombreFiltro = "Escala Grises"
			MatrizFiltro.LeerInicial("csv/"+tipo.NombreImagen+"/inicial.csv", tipo.NombreImagen)
			MatrizFiltro.EscalaGrises(tipo.NombreImagen + "Grises")
			nombreHtml = tipo.NombreImagen + "Grises"
			FiltrosColocados += "Grises, "
		case 3:
			nombreFiltro = "Espejo X"
			MatrizFiltro.LeerInicial("csv/"+tipo.NombreImagen+"/inicial.csv", tipo.NombreImagen)
			MatrizFiltro.RotacionX()
			MatrizFiltro.GenerarImagen(tipo.NombreImagen + "RX")
			nombreHtml = tipo.NombreImagen + "RX"
			FiltrosColocados += "Eje X, "
		case 4:
			nombreFiltro = "Espejo Y"
			MatrizFiltro.LeerInicial("csv/"+tipo.NombreImagen+"/inicial.csv", tipo.NombreImagen)
			MatrizFiltro.RotacionY()
			MatrizFiltro.GenerarImagen(tipo.NombreImagen + "RY")
			nombreHtml = tipo.NombreImagen + "RY"
			FiltrosColocados += "Eje Y, "
		case 5:
			nombreFiltro = "Doble Espejo"
			MatrizFiltro.LeerInicial("csv/"+tipo.NombreImagen+"/inicial.csv", tipo.NombreImagen)
			MatrizFiltro.RotacionDoble()
			MatrizFiltro.GenerarImagen(tipo.NombreImagen + "RDoble")
			nombreHtml = tipo.NombreImagen + "RDoble"
			FiltrosColocados += "Doble,  "
		}
		cmd := exec.Command("cmd", "/c", "start", "ImagenFinal/"+nombreHtml+".html")
		err := cmd.Start()
		if err != nil {
			fmt.Println(err)
		}	

		GrafoEmpleado.InsertarValores(EmpleadoLogeado, strconv.Itoa(PedidosCola.Primero.Pedido.Id_Cliente), tipo.NombreImagen, nombreFiltro)
		return c.JSON(&fiber.Map{
			"status": 200,
		})
	})

	app.Get("/obtenerPedido", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"datos": PedidosCola.Primero.Pedido,
		})
	})

	app.Post("/generarfactura", func(c *fiber.Ctx) error {
		var nuevoBloque Peticiones.BloquePeticion
		c.BodyParser(&nuevoBloque)
		FacturasRealizadas.InsertarBloque(nuevoBloque.Timestamp, nuevoBloque.Biller, nuevoBloque.Customer, nuevoBloque.Payment)
		/*Ingresar al grafo, tomar los valores de nuevoBloque.Biller, nuevoBloque.Customer, PedidosCola.Primero.Pedido.Nombre_Imagen,Filtros_colocados */
		PedidosCola.Descolar()
		VerFacturasRealizadas.NewTablaHash()
		FacturasRealizadas.InsertarTabla(VerFacturasRealizadas, EmpleadoLogeado)
		MatrizOriginal = &Matriz.Matriz{Raiz: &Matriz.NodoMatriz{PosX: -1, PosY: -1, Color: "Raiz"}}
		MatrizFiltro = &Matriz.Matriz{Raiz: &Matriz.NodoMatriz{PosX: -1, PosY: -1, Color: "Raiz"}}
		// FacturasRealizadas.GenerateGraph()
		return c.JSON(&fiber.Map{
			"datos": FacturasRealizadas.Bloques_Creados,
		})
	})

	app.Get("/facturaempleado", func(c *fiber.Ctx) error {

		return c.JSON(&fiber.Map{
			"status":  200,
			"factura": VerFacturasRealizadas.Tabla,
		})
	})

	app.Listen(":3001")
}
