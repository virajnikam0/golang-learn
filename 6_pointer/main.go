package main

import "fmt"

// changing value from pointer
func f1(param *int){
	*param = 999 // taking pointer and changing value 
} 




func main() {

	// var v1 int = 10
	// var v2 = &v1
	// *v2 = 12 // chnaging directly value from v1
	// fmt.Println(v1)
	// fmt.Println(v2)
	// fmt.Println(v1)

	// x := 100
	// f1(&x)
	// fmt.Println(x)

	// var x *int 
	// fmt.Println(x) // <nil>
	// fmt.Println(&x) // any empty address 
	// // fmt.Println(*x) // nil pointer accessing value can throw error


	// p := new(int) // this is creating new memory for int 
	// *p = 42
	// fmt.Println(*p)


	var a int = 30
	var b *int = &a
	var c *int = b

	*c = 100
	fmt.Println(*c)
	fmt.Println(*b)
	fmt.Println(a)



}