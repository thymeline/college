package main
import (
	"fmt"
)
const {
	Single = "single"
	Double = "double"
	Suite = "suite"

	Free = "free"
	Booked = "booked"
	Maintenance = "maintenance"
}

type room struct {
	typ string
	status string
	price float64
}

var rooms = map[string]room{"101": room{Single, Free, 505.5}, "102": room{Double, Maintenance, 205.5}}

func main() {
	book("101")
	book("102")
	
}

func book(num string) {
	if rooms[num].status == Free {
		rooms[num].status = Booked
		fmt.Println("Заняли комнату", num)
	} else {
		fmt.Println("Комнату", num, "сейчас нельзя занять")
	}
	fmt.Println(rooms)
}