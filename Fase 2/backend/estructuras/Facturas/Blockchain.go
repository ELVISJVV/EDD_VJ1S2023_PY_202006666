package Facturas

import (
	"backend/estructuras/TablaHash"
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"os"
	"os/exec"
	"fmt"
)

type BlockChain struct {
	Inicio          *NodoBloque
	Bloques_Creados int
}

func (b *BlockChain) InsertarBloque(fecha string, biller string, customer string, payment string) {
	cadenaFuncion := strconv.Itoa(b.Bloques_Creados) + fecha + biller + customer + payment
	hash := SHA256(cadenaFuncion)
	if b.Bloques_Creados == 0 {
		datosBloque := map[string]string{
			"index":        strconv.Itoa(b.Bloques_Creados),
			"timestamp":    fecha,
			"biller":       biller,
			"customer":     customer,
			"payment":      payment,
			"previoushash": "0000",
			"hash":         hash,
		}
		nuevoBloque := &NodoBloque{Bloque: datosBloque}
		b.Inicio = nuevoBloque
	} else {
		aux := b.Inicio
		for aux.Siguiente != nil {
			aux = aux.Siguiente
		}
		datosBloque := map[string]string{
			"index":        strconv.Itoa(b.Bloques_Creados),
			"timestamp":    fecha,
			"biller":       biller,
			"customer":     customer,
			"payment":      payment,
			"previoushash": aux.Bloque["hash"],
			"hash":         hash,
		}
		nuevoBloque := &NodoBloque{Bloque: datosBloque, Anterior: aux}
		aux.Siguiente = nuevoBloque
	}
	b.Bloques_Creados++
}

func SHA256(cadena string) string {
	hexaString := ""
	h := sha256.New()
	h.Write([]byte(cadena))
	hash := h.Sum(nil)
	hexaString = hex.EncodeToString(hash)
	return hexaString
}

func (b *BlockChain) InsertarTabla(tabla *TablaHash.TablaHash, idEmpleado string) {
	aux := b.Inicio
	for aux != nil {
		if aux.Bloque["biller"] == idEmpleado {
			tabla.Insertar(aux.Bloque["customer"], aux.Bloque["hash"])
		}
		aux = aux.Siguiente
	}
}


// func (l *BlockChain) ReturnImagen(posicion int) *Imagen{
// 	actual := l.Inicio
// 	i := 1
// 	for actual != nil{
// 		if posicion == i{
// 			return actual.Imagen
// 		}
// 		actual = actual.siguiente
// 		i += 1
// 		}
// 	 return nil
// 	}

// func (c *BlockChain) GraficarListaDoble() {
// 	nombre_archivo := "./listaDoble.dot"
// 	nombre_imagen := "listaDoble.jpg"
// 	texto := "digraph L{\n"
// 	texto += "node [shape=box fillcolor=\"#FFEDBB\" style = filled]\n"
// 	texto += "subgraph cluster_p{\n"
// 	texto += "label=\"Lista Imagenes \"\n"
// 	texto += "bgcolor = \"#398D9C\"\n"
	
	
	
	
// 	size := c.Bloques_Creados
	
// 	for i := 1; i <size +1; i++ {
		

// 		// texto += "Columna" + strconv.Itoa(i) + "[label = \""+ strconv.Itoa(c.ReturnImagen(i).GetUser())+"\n"+c.ReturnImagen(i).GetNombre() +"\", fillcolor=yellow];\n"

// 		texto += "Columna" + strconv.Itoa(i) + "[label = \""+ (c.ReturnImagen(i).Nombre)+"\n"+"\", fillcolor=yellow];\n"

// 	}


// 	texto += "{rank = same;\n"
	
// 	for i := 1; i <size +1; i++ {
// 		if i == size{
			
// 			texto += "Columna" + strconv.Itoa(i) +"}\n"
// 		}else{
// 			texto += "Columna" + strconv.Itoa(i) +";\n"
// 		}

// 	}
	
// 	for i := size; i>0 ; i-- {
// 		if i-1 == 0{
// 			break

// 		}
// 		texto += "Columna"+ strconv.Itoa(i) + "-> Columna" +strconv.Itoa(i-1)+ ";\n"
// 		texto += "Columna"+ strconv.Itoa(i-1) + "-> Columna" +strconv.Itoa(i)+ ";\n"
// 	}



// 	texto += "}\n}\n"



// 	crearArchivo(nombre_archivo)
// 	escribirArchivoDot(texto, nombre_archivo)
// 	ejecutar(nombre_imagen, nombre_archivo)
// }


// func crearArchivo(nombre_archivo string) {

// 	error := os.Remove(nombre_archivo)
// 	if error != nil {
//   		// fmt.Printf("Error eliminando archivo: %v\n", error)
// 		fmt.Printf("")
// 	} else {
//   		// fmt.Println("Eliminado correctamente")
// 		fmt.Printf("")
// 	}
// 	//Verifica que el archivo existe
// 	var _, err = os.Stat(nombre_archivo)
// 	//Crea el archivo si no existe
// 	if os.IsNotExist(err) {
// 		var file, err = os.Create(nombre_archivo)
// 		if err != nil {
// 			return
// 		}
// 		defer file.Close()
// 	}
// 	// fmt.Println("Archivo creado exitosamente", nombre_archivo)
// }

// func escribirArchivoDot(contenido string, nombre_archivo string) {
	
// 	var file, err = os.OpenFile(nombre_archivo, os.O_RDWR, 0644)
// 	if err != nil {
// 		return
// 	}
// 	defer file.Close()
// 	// Escribe algo de texto linea por linea
// 	_, err = file.WriteString(contenido)
// 	if err != nil {
// 		return
// 	}
// 	// Salva los cambios
// 	err = file.Sync()
// 	if err != nil {
// 		return
// 	}
// 	// fmt.Println("Archivo actualizado existosamente.")
// }

// func ejecutar(nombre_imagen string, archivo_dot string) {
// 	path, _ := exec.LookPath("dot")
// 	cmd, _ := exec.Command(path, "-Tjpg", archivo_dot).Output()
// 	mode := 0777
// 	_ = os.WriteFile(nombre_imagen, cmd, os.FileMode(mode))
// }


func (b *BlockChain) GenerateGraph() {
	dot := "digraph G {\n"
	dot += "\tnode [shape=box];\n"
	aux := b.Inicio
	for aux != nil {
		index := aux.Bloque["index"]
		timestamp := aux.Bloque["timestamp"]
		biller := aux.Bloque["biller"]
		customer := aux.Bloque["customer"]
		// payment := aux.Bloque["payment"]
		previousHash := aux.Bloque["previoushash"]
		// hash := aux.Bloque["hash"]

		nodeLabel := fmt.Sprintf("Timestamp: %s\nBiller: %s\nCustomer: %s\nPrevious Hash: %s",
			 timestamp, biller, customer, previousHash)

		dot += fmt.Sprintf("\t\"%s\" [label=\"%s\"];\n", index, nodeLabel)

		if aux.Siguiente != nil {
			nextIndex := aux.Siguiente.Bloque["index"]
			dot += fmt.Sprintf("\t\"%s\" -> \"%s\";\n", index, nextIndex)
		}

		aux = aux.Siguiente
	}
	dot += "}"

	// Guardar el código DOT en un archivo temporal
	file, err := os.CreateTemp("", "blockchain-graph-*.dot")
	if err != nil {
		fmt.Println("Error al crear el archivo temporal:", err)
		return
	}
	defer os.Remove(file.Name())
	defer file.Close()

	if _, err := file.WriteString(dot); err != nil {
		fmt.Println("Error al escribir en el archivo:", err)
		return
	}

	// Generar el gráfico usando Graphviz
	cmd := exec.Command("dot", "-Tpng", file.Name(), "-o", "blockchain-graph.png")
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error al generar el gráfico:", err)
		return
	}

	fmt.Println("Gráfico generado exitosamente: blockchain-graph.png")
}
