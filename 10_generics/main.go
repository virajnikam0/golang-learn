package main

import (
	"fmt"
	// "strings"
)

// generics func
// func renderingData[T interface{}](items[] T) {  // its any type generics
func renderingData[T int16 | string](items []T) {
	for _, data := range items {
		fmt.Println(data)
	}
}

// comparaeble interface
func postingData[T comparable](param T) {
	fmt.Println(param)
}

// generic struct
type user[T any] struct {
	uName T
}

func main() {
	var fruits []string
	fruits = append(fruits, "banana")
	fruits = append(fruits, "apple")
	fruits = append(fruits, "berry")

	renderingData(fruits)

	// generics structs
	user1 := user[string]{
		uName: "John",
	}

	fmt.Println(user1)

	// comparable type
	postingData("JSON object")

}
