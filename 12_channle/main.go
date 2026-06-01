package main

import "fmt"

func f1() {
	defer fmt.Println("func-1") // defer always execute in last

	fmt.Println("func-2")

}

func main() {
	// defer fmt.Println("one")
	// fmt.Println("two")
	// f1()

	bufferedChannel := make(chan int, 10)

	for i := range 5 {

	}

}
