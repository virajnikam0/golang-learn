package main

import "fmt"

// normal function
func normalFunction() string{
	return "Hello"
}
// function parameter
func paramterFunc(x1,x2,x3 int,x4,x5 string)(int,int,string){
	return x1,x2,x5 
}
// function as parameter
func helloFunc(){
	fmt.Println("Hello world")
}
func f1(op func()){
	op()
}
// function as return 
func s1() func(int)(int,string){
	return func(x int)(int,string){
		f := x 
		return f,"hey"
	}
}

// variadic function - taking no of arguments
func variadicFunction(names ... string) string {
	var name string  = ""
	for i:=0;i<len(names);i++{
		name = name + " " + names[i]
	}
	return name
}

// closure
func outer() func() int {
	var v1 int = 10
	return func() int {
		v1+=1
		return v1
	}
}







func main(){
	// f1(helloFunc)

	// x2 := s1()
	// v1,v2 := x2(1212)
	// fmt.Println(v1,v2)

	// v1,_,v3 := paramterFunc(1,2,3,"Hi","Bye")
	// fmt.Println(v1,v3)


	// names := []string{"john","ken","other"}
	// allNames := variadicFunction("john","ken","other") // passing single by single value
	// allNames := variadicFunction(names...) // passing array/splice and then converting to multiple 
	// fmt.Println(allNames)

	v1 := outer()
	v1()
	v1()
	v1()
	v1()
	fmt.Println(v1())




}


