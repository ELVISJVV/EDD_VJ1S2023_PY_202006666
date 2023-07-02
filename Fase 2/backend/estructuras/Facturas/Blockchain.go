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
