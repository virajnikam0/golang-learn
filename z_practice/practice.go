package main

import (
	"fmt"
	// "time"
)

func hey() {
	fmt.Println("hello world")
}

// john wick infomration
// name := "john wick" // cananot use outside main function 
var age int = 40
const job string = "killer"


// function

func f1(p1,p2 int,p3,p4 string)(int,string){ // function with multiple return type | if parameter have same data type we can use this way
	return p1+p2,"hey"
}


func f2(){
	fmt.Println("Hello world")
}

var funcAsVariable = f2  // function as variable


func add(x,y int)int{
	return x+y
}
func outerFunc(x,y int, innerFunct func(int,int) int) int {
	return innerFunct(x,y)
}


func m1() func(int) int {
	return func(y int)int {
			return y * y
		}
}

func clou1() func(int)int{
	var x int = 100
	return func(y int)int{
		x += y
		return x 
	}
}

func variadicFunction(p ...int) int{
	var temp int = 0  
	for _,v1 := range(p){
		temp += v1
		fmt.Println("v1: ",v1,"temp: ",temp)
	}
	return temp
}


type device struct{}
func (d device) powerOn(deviceName string){
	fmt.Println(deviceName," device power is on")
}


type Status int

const (
	Created Status = iota
	Confirmerd
	Delivered
)

func orderStatus(status Status){
	fmt.Println("order line status is: " , status)
}


func x(nums [3]int) [3]int{
	y := 0
	i := len(nums)-1
	for 0 <= i{
		y+=nums[i]
		i--
	} 
	fmt.Println(y)
	return nums
}



func main() {
	// rt1,rt2 := f1(12,12,"hey","bye")
	// fmt.Println(rt1,rt2)
	
	// funcAsVariable() // function as variable

	// funcAsReturn := outerFunc(122,12,add) // calling another fuction as through parameter
	// fmt.Println(funcAsReturn)


	// f := m1()
	// fmt.Println(f(12))

	// c1:=clou1()
	// c1(1) 
	// c1(1) 
	// temp:=c1(1)
	// fmt.Println(temp)

	// ans := variadicFunction(12,1,3,4,5,6,6,5)
	// fmt.Println(ans)


	// pointers
	// var age int = 40
	// xAge := &age
	// *xAge = 50
	// fmt.Println(*xAge)
	// fmt.Println(age)


	// orderStatus(Confirmerd)
	// orderStatus(Created)
	// orderStatus(Delivered)

	// var PI float32 = 323.2

	// switch(PI){
	// 	case 12.1,323.2:
	// 		fmt.Println("1case")
	// 	case 3.14:
	// 		fmt.Println("2case")
	// 	default:
	// 		fmt.Println("default condition")
	// }


	y := [3]int{10,10,10}
	x(y)
	
}

