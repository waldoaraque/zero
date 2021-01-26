package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	begin := time.Now()

	servers := []string{
		"http://platzi.com",
		"http://google.com",
		"http://facebook.com",
		"http://instagram.com",
	}

	for _, server := range servers {
		checkServer(server)
	}

	timePassed := time.Since(begin)

	fmt.Printf("Execution time %s\n", timePassed)
}

func checkServer(server string) {
	_, err := http.Get(server)
	if err != nil {
		fmt.Println(server, "doesn't available")
	} else {
		fmt.Println(server, "working")
	}

}