package teclado

import (
	"bufio"
	"fmt"
	"os" // package de sistema operativo
	"strconv"
)

var number1 int
var number2 int
var legend string
var err error

func EnterNumber() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Ingrese número uno: ")
	if scanner.Scan() { //si el usario ingreso algo en consola
		number1, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("El dato ingresado es incorrecto, " + err.Error())
		}
	} else {
		fmt.Println("No se ingreso algún número: ")
	}

	fmt.Println("Ingrese numero dos: ")
	if scanner.Scan() { //si el usario ingreso algo en consola
		number2, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("El dato ingresado es incorrecto, " + err.Error())
		}
	} else {
		fmt.Println("No se ingreso algún número: ")
	}

	fmt.Println("Ingresa leyenda")
	if scanner.Scan() {
		legend = scanner.Text()
	}

	fmt.Println(legend, number1*number2)

}
