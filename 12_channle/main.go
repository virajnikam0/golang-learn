package main

import (
	"fmt"
	"time"
)

func f1() {
	defer fmt.Println("func-1") // defer always execute in last

	fmt.Println("func-2")

}

func check(ch chan int, p1 int) {
	result := p1 * 2
	ch <- result
}

// -------------------------------------------
// channle producer
func producer(ch chan string) {
	ch <- "data has been received"
	close(ch)
}

// channle consumer
func consumer(ch chan string) {
	for data := range ch {
		fmt.Println(data)
	}
}

// -------------------------------------------

func main() {
	// defer fmt.Println("one")
	// fmt.Println("two")
	// f1()
	// ch := make(chan int)
	// go func() {
	// 	for i := 1; i <= 5; i++ {
	// 		ch <- i // Sending values to the channel
	// 	}
	// 	close(ch) // Closing the channel after all values are sent
	// }()

	// for x := range ch {
	// 	fmt.Println(x) // Output: 1, 2, 3, 4, 5
	// }

	// mycheck := make(chan int)
	// go check(mycheck, 10)
	// go check(mycheck, 20)
	// go check(mycheck, 30)

	// rs1 := <-mycheck
	// rs2 := <-mycheck
	// rs3 := <-mycheck

	// fmt.Println(rs1)
	// fmt.Println(rs2)
	// fmt.Println(rs3)

	ch := make(chan string)
	go producer(ch)
	go consumer(ch)

	time.Sleep(time.Second)

}
