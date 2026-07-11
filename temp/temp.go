package main

import (
	"fmt"
	"log"
	// "strconv"
	// "strings"
	// "slices"
)

func String() {
	fmt.Println("hello there")
}

func main() {
	a := 1
	if a == 1 {
		log.Fatal("yes")
	} else {
		log.Fatal("no")
	}
}
