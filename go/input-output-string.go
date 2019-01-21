package main 
import(
	"fmt"
)

func main() {
	var name string
	fmt.Println("What is your name?")
	fmt.Scanf("%v\n",&name)
	fmt.Printf("My name is %s\n", name)	
}