package main

import "fmt"

func main() {
	x := 25
	fmt.Println(x)
	y := &x // referencia la direccion en memoria
	fmt.Println(*y) // *y apunta y extrae el valor en la referencia de la memoria
}

func changeValue(a int) {
	a = 36
}