package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type calculator struct {

}

func (calculator) operate(data string, operator string) int { //receiver function asigna esta funcion como metodo a calculator
	dataClear := strings.Split(data, operator)
	op1 := parser(dataClear[0])
	op2 := parser(dataClear[1])

	switch operator {
		case "+":
			fmt.Println(op1 + op2)
			return op1 + op2
		case "-":
			fmt.Println(op1 - op2)
			return op1 - op2
		case "*": 
			fmt.Println(op1 * op2)
			return op1 * op2
		case "/":
			fmt.Println(op1 / op2)
			return op1 / op2
		default: 
			fmt.Println(operator, "No está soportado.")
			return 0
		}
}

func parser(data string) int {
	op, _ := strconv.Atoi(data)
	return op
}

func inputData() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	
	return scanner.Text()
}

func main() {
	input := inputData()
	op := inputData()
	c := calculator{}
	fmt.Println(c.operate(input, op))
}