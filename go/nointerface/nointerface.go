package main

import "fmt"

type dog struct{}

type fish struct{}

type bird struct{}

func (dog) walk() string {
	return "I'm a dog and walk..."
}

func (fish) swim() string{
	return "I'm a fish and swim..."
}

func (bird) fly() string {
	return "I'm a bird and fly..."
}

func moveDog(d dog) {
	fmt.Println(d.walk())
}

func moveFish(f fish) {
	fmt.Println(f.swim())
}

func moveBird(b bird) {
	fmt.Println(b.fly())
}

func main() {

	d := dog{}
	moveDog(d)
	
	f := fish{}
	moveFish(f)

	b := bird{}
	moveBird(b)

}