package main
import (
	"fmt"
)

func main() {
	var bron = []int{9, 10, 11, 12, 13, 14, 25, 26}
	var weekday = 0
	var cost = 0

	var result = 0

	var i int
	for i = 1; i <= 30; i++ {
		weekday++
		if weekday > 7 {weekday = 1}

		if weekday >= 1 && weekday <= 4 {
			cost = 2100
		} else {
			cost = 2850
		}
		var d int
		for _, d = range bron {
			if d == i {
				result = result + cost
			}
		}

	}
	fmt.Println(result)
}