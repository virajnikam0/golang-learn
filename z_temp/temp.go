package main

import (
	"fmt"
	// "reflect"
)

type User struct { // declaration of struct
	name  string
	age   int16
	study string
}

func (u *User) changeName(username string) {
	u.name = "john"
}

func main() {

	userTable := []User{
		User{
			name:  "John",
			age:   12,
			study: "survive",
		},
		User{
			name:  "mike",
			age:   102,
			study: "police",
		},
		User{
			name:  "white",
			age:   152,
			study: "teacher",
		},
	}

	fmt.Println(len(userTable))

	for indx := range userTable {
		fmt.Print(userTable[indx].name)
		userTable[indx].changeName(userTable[indx].name)
		fmt.Println("\n", userTable[indx].name)

		// fmt.Println(userTable[data])
	}

	// creating var through struct

	// values := reflect.ValueOf(user1)
	// // keys := values.Type()

	// for k := range values.NumField() {
	// 	data := values.Field(k)
	// 	fmt.Println(data)
	// }

}
