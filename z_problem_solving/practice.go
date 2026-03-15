package main

import "fmt"

func greet() {
	fmt.Println("Hello Ro")
	fmt.Println("I am learing Go-Lang")
}

func constVar() {
	const PI float32 = 3.14
	const APPNAME string = "Go Practice"
}

func checkIfElseCondition(num int) string {
	if num > 0 {
		return "Positive"
	} else if num < 0 {
		return "Negative"
	}
	return "Zero"
}

func loopsEvenNum() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func switchCondition(day int) {
	switch day {
	case 6, 7:
		fmt.Println("weekend")
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thrusday")
	case 5:
		fmt.Println("Friday")
	default:
		fmt.Println("not a day")
	}
}

// TODO: how to pass array as param and return array

func arrayPract(){
	// array
	// var nums = [5]int{10,20,30,40,50} 
	// nums:= []int{10,20,30,40,50}
	nums:= [5]int{10,20,30,40,50}
	fmt.Println(nums)

	sum := 0
	for i:=0;i<len(nums);i++{
		sum+=nums[i]
	}

	fmt.Println(sum)

}

func main() {
	// greet()
	// constVar()
	// fmt.Println(checkIfElseCondition(1))
	// fmt.Println(checkIfElseCondition(-1))
	// fmt.Println(checkIfElseCondition(0))
	// loopsEvenNum()
	// switchCondition(1)
	// switchCondition(6)
	// switchCondition(690)
	arrayPract()


}
