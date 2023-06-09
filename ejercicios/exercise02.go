package ejercicios

import (
	"bufio"
	"fmt"
	"os" // package de sistema operativo
	"strconv"
)

var number int
var err error

func ValidateNumber() string {
	scanner := bufio.NewScanner(os.Stdin)
	var text string
	fmt.Println("Ingrese un número entre 1 y 10: ")
	for scanner.Scan() {
		number, err = strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("El número ingresado es invalido, digite un número entre 1 y 10: ")
			continue
		} else {
			fmt.Println("El número ingresado es valido :", number)
			fmt.Println()
			for i := 1; i <= 10; i++ {
				text += fmt.Sprintf("\n %d * %d = %d", number, i, number*i)
			}
			break
		}
	}
	return text
}
