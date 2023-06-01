package variables

import (
	"fmt"
	"strconv"
	"time"
)

var Name string
var Status bool
var Salary float64
var dateTime time.Time

func OtherVariables() {
	Name = "Jesus Flores"
	Status = true
	Salary = 32242
	dateTime = time.Now()

	fmt.Println(Name)
	fmt.Println(Status)
	fmt.Println(Salary)
	fmt.Println(dateTime)

}

func ConvertToText(number int) (bool, string) {
	text := strconv.Itoa(number)
	return true, text
}
