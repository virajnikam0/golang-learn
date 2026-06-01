package main

import (
	"fmt"
	"time"
)

func f1() {
	for i := range 5 {
		fmt.Println(i)
	}
	time.Sleep(2 * time.Second)
}

func main() {

	go f1()
	f1()

}
