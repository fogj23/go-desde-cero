package interfacesimpl

import (
	"fmt"

	"github.com/go-desde-cero/interfaces"
)

func HumanoRespirando(ih interfaces.Humano) {
	ih.Respirar()
	fmt.Printf("Soy un/a %s, y estoy respirando \n", ih.Genero())
}
