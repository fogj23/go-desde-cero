package iteraciones

import (
	"fmt"
)

func Iterated() {
	for i := 0; i < 12; i += 2 {
		fmt.Println(i)
	}

	for b := 0; b < 12; b++ {
		fmt.Println(b)
	}

	for b := 0; b < 12; b++ {
		if b == 6 {
			continue
		}
		fmt.Println(b)
	}
}
