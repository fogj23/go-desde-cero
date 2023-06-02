package files

import (
	"bufio"
	"fmt"
	"os"

	"github.com/go-desde-cero/ejercicios"
)

var filename string = "./files/txt/table.txt"

func SaveTable() {
	var text = ejercicios.ValidateNumber()
	file, err := os.Create(filename)

	if err != nil {
		fmt.Println("Error al crear archivo " + err.Error())
		return
	}
	fmt.Fprintln(file, text)
	file.Close()
}

func AddTable() {
	var text string = ejercicios.ValidateNumber()
	if Append(filename, text) {
		fmt.Println("Error al concatenar contenido")
	}
}

func Append(filen string, text string) bool {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		fmt.Printf("Error durante Append" + err.Error())
		return false
	}

	_, err = file.WriteString(text)
	if err != nil {
		fmt.Printf("Error durante WriteString" + err.Error())
		return false
	}

	file.Close()
	return true
}

func ReadFile() {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error al leer archivo" + err.Error())
		return
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		item := scanner.Text()
		fmt.Println(">>> " + item)
	}
	file.Close()
}
