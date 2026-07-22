package main

import "fmt"

func FindSencondLastBigNum() {
	var v1, v2, v3 int = 0, 0, 0
	fmt.Scan(&v1)
	fmt.Scan(&v2)
	fmt.Scan(&v3)

	if v1 > v2 && v1 < v3 {
		fmt.Println(v1)
	} else if v2 > v1 && v2 < v3 {
		fmt.Println(v2)
	} else if v3 > v1 && v3 < v2 {
		fmt.Println(v3)
	}

}

func ChoklateCost() {
	var a, b, c int = 4, 4, 1000
	a = a * 5
	b = b * 10
	noofChocklateToBuy := (a + b) / c
	fmt.Println(noofChocklateToBuy)
}

func main() {
	// FindSencondLastBigNum()
	// ChoklateCost()

	var a, b int = 0, 0
	a = 2
	b = 3
	c := (b-a)*2 + a
	fmt.Println(c)

}
