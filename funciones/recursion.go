package funciones

import "fmt"

func Exponencia(valor int) {
	fmt.Println("exponencia=", valor)
	if valor >= 100 {
		return
	}
	Exponencia(valor * 2)
}
