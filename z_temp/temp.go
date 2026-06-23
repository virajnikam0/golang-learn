package main

import (
	"fmt"
	"math/rand"
)

type person struct {
	name string
	age  int
}

// global array
var persons []*person

func f1() {
	for i := range 3 {
		personName := fmt.Sprintf("%d person", i)
		p := person{
			name: personName,
			age:  i,
		}
		persons = append(persons, &p)
	}
}

func f2() {
	for i := 0; i < len(persons); i++ {
		var pp person
		pp = *persons[i]
		fmt.Println(pp.name)
	}
}

func main() {
	// f1()
	// f2()

	for i := range 5 {
		status := []string{
			"a",
			"b",
			"c",
		}[rand.Intn(3)]
		fmt.Println(i)
		fmt.Println(status)
	}

}
