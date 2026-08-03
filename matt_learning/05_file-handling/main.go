package main

import (
	"fmt"
	"os"
)

func main() {
	// create file

	file, err := os.Create("data.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer file.Close()
	// write file
	data := "hello there"
	data2 := "nice to meet you"

	n, err := file.WriteString(data)
	n1, err := file.WriteString(data2)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(n)
		fmt.Println(n1)
	}

	// read file
	rdfile, err := os.Open("data.txt")

	buff := make([]byte, 1024)

	n11, err := rdfile.Read(buff)

	if err != nil {

	}

	fmt.Println(string(buff[:n11]))

}
