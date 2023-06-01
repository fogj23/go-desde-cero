package variables

import (
	"fmt"
)

func ShowIntegers() {
	var intcomun int
	intDe32 := int32(10)
	intDe64 := int64(100)

	fmt.Println("intcomun: ", intcomun)
	fmt.Println("intDe32: ", intDe32)
	fmt.Println("intDe64: ", intDe64)
}
