package abx

import "fmt"

func H1() {
	fmt.Println("hello there")
}

func H2() {
	fmt.Println("hello there")
}

func PoinerrFunc() {
	var v1 int = 116
	var v2 *int = &v1

	*v2 = 229

	fmt.Println(*v2)
	fmt.Println(v1)

}
