package main
import (
	"fmt"
)

type employee struct {
	id int;
	name string;
	position string;
	salary float64
}

func main() {
	var employees []employee = []employee{ employee{id: 0, name: "алексей", position: "директор", salary: 15.5}, employee{id: 1, name: "никита", position: "раб", salary: 7.2}, employee{id: 2, name: "андрей", position: "раб", salary: 3.2} }
	calc(employees)
}

func calc(employees []employee) {
	var check employee

	var sum float64 = 0.0
	var dot float64
	for _, check = range employees {
		sum = sum + check.salary
	}
	dot = sum / float64(len(employees))
	
	fmt.Println("сумма", sum)
	fmt.Println("среднее", dot)
}