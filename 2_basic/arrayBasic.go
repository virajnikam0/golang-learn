package main

import (
	"fmt"
	"slices"
)

func main() {

	
	// array 
	var nums[4]int // df value = 0
	var names[4]string // df value = ''
	var isTrue[4]bool // df value = false

	nums[1] = 223
	names[1] = "john"
	isTrue[1] = true

	fmt.Println(nums)
	fmt.Println(names)
	fmt.Println(isTrue)

	fmt.Println(len(nums))
	fmt.Println(len(names))
	fmt.Println(len(isTrue))

	
	x := [3]int{12,12,334}
	fmt.Println(x)


	// matrix array

	matrix := [3][3]int{
		{1,2,3},
		{4,5,6},
		{7,8,9},
	}
	fmt.Println(matrix)

	// looping array
	for i :=0; i<len(x);i++{
		fmt.Println(x[i])
	}

	// looping matrix
	val := 0
	for i:=0;i<len(matrix);i++{
		for j:=0;j<len(matrix[i]);j++{
			val = val + 1
			fmt.Print(matrix[i][j] , " value is: " )
			fmt.Println(val)

		}
	}


	// slice 
	var s1[] int
	fmt.Println(len(s1)) // nil 
	
	s3 := make([]int,3)
	s3[0] = 100
	s3[1] = 200
	s3[2] = 300
	// s3[4] = 400
	fmt.Println(s3)

	s4 := make([]int,2,5) // 2 parameter adding value in front of array [0,0,101,202]
	// s4 := make([]int,0,5) //  adding 0 in 2nd parameter can [101,202]
	s4 = append(s4,101)
	s4 = append(s4,101)
	s4 = append(s4,101)
	// s4 = append(s4,101)
	println("----------------------")
	fmt.Println(s4)
	fmt.Println(cap(s4))
	
	s5 := []int{} // empty splice size=0 capacity=0
	println("----------------------")
	fmt.Println(len(s5)) 
	fmt.Println(cap(s5)) 


	//==============================================

	var s6 []int
	s6 = append(s6,10)
	s6 = append(s6,20)
	s6 = append(s6,30)

	s7 := make([]int,len(s6)) // allocate space

	copy(s7,s6) // copy data from sourse -> destination

	fmt.Println(s7)
	fmt.Println(s6)

	fmt.Println("slice opearator")
	fmt.Println(s6[0:])
	// fmt.Println(s6[:5]) // runtime error cause array is out of bound
	fmt.Println(s6[:3]) // equal to length it wont throw error
	fmt.Println(s6[:])

	// slices methods 

	s11 := []int{}
	s11 = append(s11,122)
	s11 = append(s11,121)
	s11 = append(s11,123)

	s12 := []int{}
	s12 = append(s12,121)
	s12 = append(s12,122)
	s12 = append(s12,123)

	fmt.Println(slices.Equal(s11,s12))

	






	



 




}