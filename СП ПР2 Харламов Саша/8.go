package main
import (
	"fmt"
	"strconv"
)

type entry struct {
	ip string;
	http int;
	td string
}

func main() {
	var entries []entry = []entry{entry{ip: "1.01.02.1", http: 405, td: "12:05:00"}, entry{ip: "1.01.02.1", http: 505, td: "12:05:01"}, entry{ip: "1.01.02.2", http: 502, td: "12:05:03"}}
	sort(entries)
}

func sort(entries []entry) {
	var user []entry
	var server []entry

	var log entry
	for _, log = range entries {
		var num = log.http
		var numstr string

		numstr = strconv.Itoa(num)
		if string(numstr[0]) == "4" {
        	user = append(user, log)
    	} else if string(numstr[0]) == "5" {
        	server = append(server, log)
    	}
	}
	fmt.Println("клиентные:", user)
	fmt.Println("серверные:", server)
}