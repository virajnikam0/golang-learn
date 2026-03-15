package main

import (
	"fmt"
	"time"
)

type Person struct {
	name   string
	age    int
	time_t time.Time
}
// struct function
func (p *Person) chnageAge(age int){
	p.age = age 
	fmt.Println("person age chnage")
}

// struct constructor
func newPerson(name string,age int,time_t time.Time) *Person{
	myPerson := Person{
		name: name,
		age: age,
		time_t: time_t,
	}

	return &myPerson

}


func main() {

	// var john Person
	// john.age = 12
	// john.name = "john wick"
	// john.time_t = time.Now()

	// fmt.Println(john)
	
	// john.chnageAge(13)
	
	// fmt.Println(john)

	// ken := newPerson("ken",12,time.Now())
	// fmt.Println(ken)

	// nameless contructor 
	killer1 := struct{
		name string
		age int
		isGood bool
	}{"john wick",42,true}

	fmt.Println(killer1)

	




}
