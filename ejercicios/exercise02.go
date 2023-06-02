package ejercicios

import (
	"bufio"
	"fmt"
	"os" // package de sistema operativo
	"strconv"
)

var number int
var err error

func ValidateNumber() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Ingrese un número entre 1 y 10: ")
	for scanner.Scan() {
		number, err = strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("El número ingresado es invalido, digite un número entre 1 y 10: ")
			continue
		} else {
			fmt.Println("El número ingresado es valido :", number)
			textNumber := strconv.Itoa(number)
			for i := 1; i <= 10; i++ {
				textI := strconv.Itoa(i)
				fmt.Println(textNumber+" * "+textI+"=", number*i)
			}
			break
		}
	}
	fmt.Println("end function")
}
