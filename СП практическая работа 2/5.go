package main
import (
	"fmt"
	"strings"
	"errors"
)

func main() {
	var x string
	var y int
	var z string

	fmt.Println("наме:")
	fmt.Scan(&x)
	fmt.Println("аге:")
	fmt.Scan(&y)
	fmt.Println("емаил:")
	fmt.Scan(&z)

	fmt.Println(validate_user(x, y, z))
}

func validate_user(name string, age int, email string) error {
	if !(len(name) > 0 && len(name) < 50) {return errors.New("Имя должно состоять из не более чем 50 символов")}
	if !(age >= 18 && age <= 120) {return errors.New("Возраст должен быть выше 18 и ниже 120")}
	if !(strings.Contains(email, "@")) {return errors.New("Некорректная почта")}
	return nil
}