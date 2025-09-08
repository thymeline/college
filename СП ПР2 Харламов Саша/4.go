package main
import (
	"fmt"
)
var kandidati = map[string]int{"Анна": 0, "Борис": 0, "Виктортор": 0}
var golosa = map[string]string{}

func main() {
	var x string
	var y string

	var name string
	for {
		fmt.Println("кто голосует?")
		fmt.Scan(&x)
		fmt.Println("за кого голосует", x, "?")
		for name, _ = range kandidati {
			fmt.Println("кандидать:", name)
		}
		fmt.Scan(&y)
		golos(x, y)
	}

}
func golos(name string, kandidat string) {
	var isUsed bool
	_, isUsed = golosa[name]
	if isUsed {return}


	var check string
	var checknum int
	var tof = false
	for check, checknum = range kandidati {
		if check == kandidat {
			tof = true 
			kandidati[check] = checknum + 1
		} 
	}

	if tof == false {
		fmt.Println("За таких мы не голосуем здесь")
		return
	}
	golosa[name] = kandidat

	fmt.Println(golosa)
	newfunc()
	
}

func newfunc() {
	var x string
	var y int

	var all int = 0
	for _, y = range kandidati {
		all += y

	}

	for x, y = range kandidati {
		if all > 0 {
			var percent = (float64(y) / float64(all)) * 100
			fmt.Println("За", x, "проголосовало", y, "человек. Это", percent, "% от всех.")
		} else {
			fmt.Println("За", x, "проголосовало", y, "человек. Это 0% от всех.")
		}
	}
}

