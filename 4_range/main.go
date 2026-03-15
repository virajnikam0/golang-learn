package main

import "fmt"

func main() {

	nums := []int{1, 2, 3, 4}

	for index, value := range nums { // range return 1st index 2nd value
		fmt.Println(index)
		fmt.Println(value)
	}
	fmt.Println("-----------------")

	for i,j := range "🤖👽👾"{  // return index so index is varry between number of unicode value here each char is takeing 4 byte so its 
		fmt.Println(i)
		fmt.Println(j)
	}
	// 0 = index
	// 129302 = unicode
	// 4 = index
	// 128125 = unicode
	// 8 = index
	// 128126 = unicode

	

}
