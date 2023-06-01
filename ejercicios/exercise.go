package ejercicios

import (
	"fmt"
	"strconv"
)

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
}
