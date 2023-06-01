package ejercicios

import (
	"fmt"
	"strconv"
)

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
}
