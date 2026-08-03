package main

import "fmt"

func h1() func() int {
	a, b := 0, 1
	return func() int {
		fmt.Println(&a, " ", &b)
		a, b = b, a+b
		return b
	}
}

func main() {

	rtValue := h1()

	for x := rtValue(); x < 100; x = rtValue() {
		fmt.Println(x)
	}

}
