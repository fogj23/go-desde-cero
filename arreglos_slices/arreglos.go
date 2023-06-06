package arreglos_slices

import "fmt"

var table [10]int
var table2 [10]int = [10]int{1, 2, 3, 4}
var matrix [3][5]int

func ShowArray() {

	table[2] = 2
	table[9] = 12

	table3 := [5]string{"Jesus", "Eleazar", ""}

	matrix[2][4] = 2

	fmt.Println(":::showArray1:::")
	fmt.Println(table)
	fmt.Println()
	fmt.Println(":::showArray2:::")
	fmt.Println(table2)
	fmt.Println()
	fmt.Println(":::showArray3:::")
	for i := 0; i < len(table3); i++ {
		fmt.Println(table3[i])
	}
	fmt.Println()
	fmt.Println(":::showMatrix:::")
	fmt.Println(matrix)
	fmt.Println()

}
