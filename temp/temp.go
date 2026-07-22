package main

import (
	// "dsa-practice/temp/abx"
	// goroutines "dsa-practice/temp/go-routines"
	"fmt"
	"sync"
	// "time"
	// "log"
	// "strconv"
	// "strings"
	// "slices"
)

var wg sync.WaitGroup

func String() {
	fmt.Println("hello there")
}

func worker(param int) {
	defer wg.Done()
	fmt.Println(param)
}

func main() {
	var x int
	fmt.Scan(&x)
	fmt.Println(x)

}
