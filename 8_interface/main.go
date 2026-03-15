package main

import "fmt"

// interface  - what object have and can be do by other after implementation of interface
type Electonics interface {
	display() string
}

type Mobile struct {}
func (m Mobile) display() string{
	return "my mobile display"
}

type Laptop struct{}
func (l Laptop) display() string{
	return "my laptop display"
}








func main(){

	var device Electonics

	device = Laptop{}
	fmt.Println(device.display())
	device = Mobile{}
	fmt.Println(device.display())




}