package main

import "fmt"

type Person struct {
	name string
}

func (P *Person) getInfo(pname string) {
	P.name = pname
}

func main() {

	p1 := Person{}

	p1.getInfo("John")

	fmt.Println(p1.name)

}
