package main
import (
	"fmt"
)
var id = 0
var card = make(map[int]Order)


type Order struct {
	items string
	total int
	address string
	isCompleted bool
}

func main() {
	var a string
	var b int
	var c string

	for {
		fmt.Println("предметы заказа:")
		fmt.Scanln(&a)
		fmt.Println("тотал заказа:")
		fmt.Scanln(&b)
		fmt.Println("адрес:")
		fmt.Scanln(&c)

		make_order(a, b, c)
	}
}

func make_order(items string, total int, address string) {
	var new_order = Order{
		items: items,
		total: total,
		address: address,
		isCompleted: false,
	}
	card[id] = new_order
	id++
	fmt.Println(card)
}
