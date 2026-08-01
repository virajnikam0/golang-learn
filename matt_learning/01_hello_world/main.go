package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// taking user input and greet him/her
	reader := bufio.NewReader(os.Stdin)
	const word string = "nice"
	data, _ := reader.ReadString('\n')
	fmt.Println("Hello Go Developer ", data)
}
