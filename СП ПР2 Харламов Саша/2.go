package main
import (
	"fmt"
)

func main() {
	fmt.Println("оСНОВНОЙ БАГАЖ:")
	var x float64
	fmt.Scanln(&x)
	fmt.Println("рУЧНАЯ КЛАДЬ:")
	var y float64
	fmt.Scanln(&y)
	fmt.Println("доп. ручная КЛАДЬ:")
	var z float64
	fmt.Scanln(&z)

	var result float64 = x+y+z
	fmt.Println("вес", result)
}
