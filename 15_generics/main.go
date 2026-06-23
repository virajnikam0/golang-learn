package main

import "fmt"

func doTheOperation[T int](v1, v2 T, opear string) T {
	switch opear {
	case "+":
		return any(v1 + v2).(T)
	case "-":
		return any(v1 - v2).(T)
	}
	return -1
}

func genericArray[T string]() []T {
	var fruits = []T{
		"banana", "mango", "grapes",
	}
	return fruits
}

func main() {
	rtData := doTheOperation(12, 12, "*")
	fmt.Println(rtData)
}
