package ejercicios

import (
	"fmt"
	"strconv"
)

<<<<<<< HEAD
func ValidateInteger(input string) (int, string) {
	fmt.Println("::::validateInteger::::")

	var message string

	number, err := strconv.Atoi(input)

	if err != nil {
		return 0, "Error: " + err.Error()
	}

	if number > 100 {
		message = "Es mayor a 100 "
		fmt.Println("error: ", err)
	} else {
		message = "Es menor a 100 "
		fmt.Println("error: ", err)
	}

	return number, message
=======
func ValidateInteger(input string) (int64, string) {
	fmt.Println("::::validateInteger::::")

	var message string
	var num int64

	if number, err := strconv.ParseInt(input, 10, 64); number > 100 {
		message = "Es mayor a 100 "
		fmt.Println(message, err)
		num = number
	} else {
		message = "Es menor a 100 "
		fmt.Println(message, err)
		num = number
	}

	return num, message
>>>>>>> 7923f1dc98dc6aff148f045c59bb14674909f8f8
}
