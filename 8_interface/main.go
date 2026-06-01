package main

import "fmt"

// interface => 2 childs =>

type Parent interface {
	height() string
	color() string
}

type Child1 struct {
	cname string
}
type Child2 struct {
	cname string
}

func (C *Child1) height() string {
	return "child 1 height"
}

func (C *Child1) color() string {
	return "child 1 color"
}
func (C *Child2) height() string {
	return "child 2 height"
}

func (C *Child2) color() string {
	return "child 2 color"
}

func main() {

	// interface var
	var p1 Parent

	p1 = &Child1{
		cname: "john wick",
	}

	fmt.Println(p1.color())

	p1 = &Child2{
		cname: "Ken thomson",
	}
	fmt.Println(p1.color())

}
