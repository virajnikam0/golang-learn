package main

import (
	"fmt"
	"time"
)


const ( 
	name = "john wick"
	age = 12
	job = "killer"
)

func main() {
	fmt.Println("hello world")

	// variable delcaration
	// 1
	var age int = 10
	fmt.Println(age)
	// 2
	var name = "john"
	fmt.Println(name)
	// 3
	study := "Biotechnology"
	fmt.Println(study)


	// constant delcaration 
	const PI float32 = 3.14
	fmt.Println(PI)
	// PI = 121.1 // cannot assign after defination
	fmt.Println(name,age,job);
	

	// loop
	// there is only for loop avaliable 

	n1 := 10 
	for n1 > 0{   // while loop  
		fmt.Println(n1)
		n1--
	}

	// for{    // infinite loop
	// 	fmt.Println("1") 
	// }
	fmt.Println("---------------------------") 
	
	for i:=0; i<10; i++ {
		if i % 2 == 0{
			continue
			}else if i == 7{
				break
		}else{
			fmt.Println(i)
		}
	}
	fmt.Println("---------------------------") 
	
	for i:= range 3 { // range functionality
		fmt.Println(i)
	}


	// if-else condition
	i := 18
	haveVoterID := true

	if i > 18 || haveVoterID {
		fmt.Println("mature age")
	}else if i == 18 && haveVoterID {
		fmt.Println("younng age")
	}else {
		fmt.Println("still boy")
	}

	fmt.Println("---------------------------") 

	// switch condition
	switch time.Now().Weekday(){    // multi-conditonal switch
		case time.Saturday, time.Sunday:
			fmt.Println("its weekend") 
		default :
			fmt.Println("its working day" + time.Now().Weekday().String())
	}


	whoAmI := func(i interface{}){
		switch i.(type){
			case int:
				fmt.Println("interger")
			case string:
				fmt.Println("string")
			case bool:
				fmt.Println("boolean")
			default:
				fmt.Println("other")
		}
	}

	whoAmI("go")
	whoAmI(121)
	whoAmI(true)
	whoAmI(121.1)


	sayHello := func(){
		fmt.Println("Hello world")
	}

	sayHello()









}
