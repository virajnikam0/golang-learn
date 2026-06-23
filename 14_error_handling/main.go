package main

import (
	"errors"
	"fmt"
)

func divideByZero() {
	zero := 0
	ans := 12 / zero
	fmt.Println(ans)

}

func deferFunction() {
	r := recover()
	fmt.Println(r, "nice to meet you")
}

// custom error function
func sayingHelloFunction() error {
	return errors.New("this is my error as saying hello error")
}

func main() {
	defer deferFunction()
	// divideByZero()
	err := sayingHelloFunction()
	fmt.Println(err)

	fmt.Println("nice")
}
