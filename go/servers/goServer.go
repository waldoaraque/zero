package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	begin := time.Now()

	channel := make(chan string)

	servers := []string{
		"http://platzi.com",
		"http://google.com",
		"http://facebook.com",
		"http://instagram.com",
	}

	i := 0
	for { //while 
		if i > 10 {
			break
		}
		for _, server := range servers {
			go checkServer(server, channel) //ejecución asíncrona por hilo de ejecución
		}
		time.Sleep(1*time.Second)
		fmt.Println(<-channel)
		i++
	}

	/*
	for i := 0; i < len(servers); i++ {
		fmt.Println(<-channel)	
	}
	*/
	timePassed := time.Since(begin)

	fmt.Printf("Execution time %s\n", timePassed)
}

func checkServer(server string, channel chan string) {
	_, err := http.Get(server)
	if err != nil {
		//fmt.Println(server, "doesn't available")
		channel <- server + " doesn't available."
	} else {
		//fmt.Println(server, "working")
		channel <- server + " working."
	}

}