package main

import (
	// "fmt"
	"matt_learning/02_hello_again/o3"
	"os"
	// "os"
)

func main() {

	// arr := os.Args
	
	if len(os.Args) > 1{
		o3.SayHello(os.Args[1])
	}else{
		o3.SayHello("Test")
	}
}
