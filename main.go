package main

import (
	"fmt"

	"github.com/go-desde-cero/ejercicios"
	"github.com/go-desde-cero/variables"

	"runtime"
)

func main() {

	fmt.Println(":::::::::::::::::ShowIntegers:::::::::::::::::::::")
	variables.ShowIntegers()
	/*
		fmt.Println(":::::::::::::::::OtherVariables:::::::::::::::::::::")
		variables.OtherVariables()
		fmt.Println(":::::::::::::::::ConvertToText:::::::::::::::::::::")
		status, text := variables.ConvertToText(10324)

		fmt.Println("status: ", status)
		fmt.Println("text: ", text)
	*/
	fmt.Println(":::::::::::::::::Condiciones:::::::::::::::::::::")

	if os := runtime.GOOS; os == "windows" {
		fmt.Println("SI esto es:", os)
	} else {
		fmt.Println("DE OTRO MODO esto Linux")
	}

	switch os := runtime.GOOS; os {
	case "windows":
		fmt.Println("case windows")
	case "mac":
		fmt.Println("case mac")
	default:
		fmt.Println("default case")
	}

	fmt.Println(":::::::::::::::::Ejercicio01:::::::::::::::::::::")

	value, message := ejercicios.ValidateInteger("101")

	fmt.Println("response number:", value)
	fmt.Println("response message:", message)

}
