package main

import (
	"fmt"
	"net/http"
	//"log"
	"io"
)

type webWriter struct {}

func (webWriter) Write(p []byte) (int, error) {
	fmt.Println(string(p)) //convertimos el slice de bytes a string
	return len(p), nil //len(p) devuelve todo el string
}

func main() {
	res, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println(err)
	}

	e := webWriter {}

	io.Copy(e, res.Body) //escritor, lector - Copy() tiene una interface implicita
	
}