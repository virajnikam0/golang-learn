package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// var str string = "hello there"
	// str += "nice to meet you"
	// str[0] = 's'

	// for _, data := range *left {
	// 	fmt.Println(data)
	// }
	// fmt.Println(str)

	// fmt.Println(&str)
	// fmt.Println(str)
	// str = strings.ToUpper(str)
	// fmt.Println(&str)
	// fmt.Println(str)

	// scan := bufio.NewScanner(os.Stdin)
	// scan.Text()

	// s1 := "heyy  there there there  nice to"
	// s1 = strings.Replace(s1, "there", "yoo ", 1)
	// fmt.Println(s1)

	// ---------------------------------------------------

	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "not enough input data ")
		os.Exit(-1)
	}

	oldname, newname := os.Args[1], os.Args[2]

	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])
	fmt.Println(os.Args[2])

	scan := bufio.NewScanner(os.Stdin)

	for scan.Scan() {
		line := scan.Text()
		str := strings.Replace(line, oldname, newname, -1)
		fmt.Println(str)

	}

	if err := scan.Err(); err != nil {
		fmt.Println(err)
	}

	// fmt.Println(scan.Text())

	// myNewString := strings.Replace(scan.Text(), oldname, newname, -1)

	// fmt.Println(myNewString)

}
