package main

import (
	"backend/estructuras/Lista"
	"backend/estructuras/Peticiones"
	"backend/estructuras/ArbolAVL"
	"fmt"
	"io/ioutil"
	"encoding/base64"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var ListaEmpleados *Lista.ListaSimple
var ArbolPedidos *ArbolAVL.Arbol

func main() {
	// estructuras utilizadas
	ListaEmpleados = &Lista.ListaSimple{Inicio: nil, Longitud: 0}
	ArbolPedidos = &ArbolAVL.Arbol{Raiz: nil}

	app := fiber.New()
	app.Use(cors.New())

	app.Post("/login", func(c *fiber.Ctx) error {
		var usuario Peticiones.Login
		c.BodyParser(&usuario)
		fmt.Println(usuario)

		if usuario.Username == "ADMIN_202006666" && usuario.Password == "admin" {
			return c.JSON(&fiber.Map{
				"status": "admin",
			})
		}else {
			if ListaEmpleados.Inicio != nil {
				if ListaEmpleados.Buscar(usuario.Username, usuario.Password) {
					// VerFacturasRealizadas = &TablaHash.TablaHash{Capacidad: 5, Utilizacion: 0}
					// VerFacturasRealizadas.NewTablaHash()
					// EmpleadoLogeado = usuario.Username
					return c.JSON(&fiber.Map{
						"status": "employee",
					})
				}
			}else {
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
		ArbolPedidos.Graficar()
		// fmt.Println(pedidosNuevos)


		// ArbolPedidos.RecorridoInorden(ArbolPedidos.Raiz, PedidosCola)
		return c.JSON(&fiber.Map{
			"status": 200,
			// "cola":   PedidosCola,
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
	app.Listen(":3001")
}
