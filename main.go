package main

import ( //"runtime"

	e "github.com/go-desde-cero/interfacesImpl"
	"github.com/go-desde-cero/modelos"
)

func main() {
	/*
			fmt.Println(":::::::::::::::::ShowIntegers:::::::::::::::::::::")
			variables.ShowIntegers()759658

				fmt.Println(":::::::::::::::::OtherVariables:::::::::::::::::::::")
				variables.OtherVariables()
				fmt.Println(":::::::::::::::::ConvertToText:::::::::::::::::::::")
				status, text := variables.ConvertToText(10324)

				fmt.Println("status: ", status)
				fmt.Println("text: ", text)

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

			value, message := ejercicios.ValidateInteger("10i")

			fmt.Println("response number:", value)
			fmt.Println("response message:", message)



		fmt.Println(":::::::::::::::::Teclado:::::::::::::::::::::")
		teclado.EnterNumber()

		iteraciones.Iterated()
	*/

	//fmt.Println(ejercicios.ValidateNumber())
	//files.SaveTable()
	//files.AddTable()
	//files.ReadFile()
	//funciones.Exponencia(2)

	//arreglosslices.ShowArray()
	//arreglosslices.Capacity()
	//mapas.ShowMaps()
	//usuarios.SaveUser()

	pedro := new(modelos.Hombre)
	e.HumanoRespirando(pedro)

	maria := new(modelos.Mujer)
	e.HumanoRespirando(maria)
}
