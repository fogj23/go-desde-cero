package usuarios

import (
	"fmt"
	"time"

	"github.com/go-desde-cero/modelos"
)

func SaveUser() {
	user := new(modelos.User)
	user.AddUser(123, "Eleazar", time.Now(), true)
	fmt.Println(user)
}
