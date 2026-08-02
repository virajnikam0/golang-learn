package main

import "fmt"

func main() {

	// var (
	// 	a = 10
	// 	// b = 20
	// )

	// fmt.Printf("%8T %8T %[1]v %[2]s", a, "bebe")
	// fmt.Printf("%8T %[1]v", b)

	// converting string to byte or rune
	str := "abc"
	arr1 := []rune(str)
	arr2 := []byte(str)

	fmt.Printf("%8T %[1]v", arr1)
	fmt.Printf("%8T %[1]s", arr2)

}
