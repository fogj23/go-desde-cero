package main

import (
	"fmt"

	"github.com/go-desde-cero/variables"
)

func main() {
	fmt.Println(":::::::::::::::::ShowIntegers:::::::::::::::::::::")
	variables.ShowIntegers()
	fmt.Println(":::::::::::::::::OtherVariables:::::::::::::::::::::")
	variables.OtherVariables()
	fmt.Println(":::::::::::::::::ConvertToText:::::::::::::::::::::")
	status, text := variables.ConvertToText(10324)

	fmt.Println("status: ", status)
	fmt.Println("text: ", text)
}
