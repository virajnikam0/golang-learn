package main

import (
	"fmt"
	"time"
)

func countNum(i int) {
	fmt.Println(i)
}

// channled passed functio
func channledFunc(ch *chan int) {
	fmt.Println("my channle value :", <-*ch)
	*ch <- 200

}

// sum func for reciving channle data
func sumFunc(cln chan int, v1, v2 int) (ch chan int) {

	cln <- v1 + v2
	return cln
}

func main() {

	// channle are created due to passing data from one routines to another routines

	chanl := make(chan int)
	chanl2 := make(chan int)

	go channledFunc(&chanl)
	chanl <- 101

	// for i := 0; i < 10; i++ {
	// 	go countNum(i)
	// }

	fmt.Println("my channle value :", <-chanl)
	time.Sleep(time.Second * 2)

	// reciving channle data
	go sumFunc(chanl2, 10, 20)
	x := <-chanl2

	fmt.Println(x)

}
