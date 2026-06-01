package main

import "fmt"

// custoum variable
type order int

const (
	SalesOrder order = iota
	CustmouOrder
	ManageOrder
)

type Gender string

const (
	M = "Male"
	F = "Female"
	O = "Nothing"
)

func main() {
	var x order = ManageOrder
	fmt.Println(x)

	var gender Gender = M
	fmt.Println(gender)

}
