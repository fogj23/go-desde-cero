package arreglos_slices

import "fmt"

var tablas []int = []int{1, 5, 4}
var array [10]int = [10]int{12, 44, 65, 2, 23, 56, 78, 97}

func ShowSlice() {
	fmt.Println("tablas: ", tablas)

	positions := array[3:]   //Slice creado desde un vector, desde la posición 3
	positions2 := array[:5]  //Slice creado desde un vector, desde la posición 0 hasta la 5
	positions3 := array[6:7] //Slice creado desde un vector, desde la posición 6 hasta la 7

	fmt.Println(positions)
	fmt.Println(positions2)
	fmt.Println(positions3)
}

func Capacity() {
	elements := make([]int, 5, 20)

	fmt.Printf("Length: %d, Capacity: %d", len(elements), cap(elements))

	list := make([]int, 0, 0)

	for i := 0; i < 14; i++ {
		list = append(list, i)
	}
	fmt.Printf("\n Length: %d, Capacity: %d", len(list), cap(list))
}
