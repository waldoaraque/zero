package main //package name
// package	otroPaquete 

import "fmt" // librerías necesarias para importar el paquete

func main() {
	// Variables
	var mensaje string = "Hola mundo"
	easyMessage := "Hola mundo usando :="
	fmt.Println(mensaje)
	fmt.Println(easyMessage)

	// Puedes comentar usando "//"

	// float numbers
	a := 10.
	var b float64 = 3
	fmt.Println(a / b)


	//integer numbers
	var c int = 10
	d := 3
	fmt.Println(c / d)

	// boolean
	var x bool = true
	y := false
	fmt.Println(x || y)
	fmt.Println(x && y)
	fmt.Println(!x)

	//Lógica privada
	privada()

	//Lógica Publica
	Publica()

	//Función "defer"
	printHelloWorld()

	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}

	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}
}

// func types
func privada() {
	fmt.Println("Ejecutar lógica que no necesita ser exportada (pertenece solo a este paquete)")
}

func Publica() {
	fmt.Println("Lógica que quiero exportar a otros paquetes")
}

// defer -> paso final en ejecutar... Espera que ejecute todo.
func printHelloWorld() {
	defer fmt.Println("World!")
	fmt.Println("Hello")

	fmt.Println("defer ejecuta un fragmento de código hasta que la función haya terminado")
}
