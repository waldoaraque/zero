package main
import(
	"fmt"
	"strconv"
)

func main() {
	edad := "22"
	edad_int,_ := strconv.Atoi(edad) //method to convert string to int
	// ,_ discard the variable that receives
	fmt.Println(edad_int)
}