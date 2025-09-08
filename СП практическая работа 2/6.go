package main
import (
	"fmt"
)

func main() {
	var posts [][]string = [][]string{{"go", "backend"}, {"git", "go", "tools"}}
	fmt.Println(check(posts))
}

func check(list [][]string) map[string]bool {
	quantity := make(map[string]int)
	result := make(map[string]bool)

	var slice []string
	var value string
	for _, slice = range list {
		for _, value = range slice {
			quantity[value] += 1
		}
	}

	var num int
	for value, num = range quantity {
		result[value] = (num == 1)
	}

	return result
}