package main

import (
	"fmt"
	"maps"
)

func main() {

	// maps => associative data structure
	// key : value

	nums := make(map[int]string)
	nums[1] = "one" // setting value in map
	nums[2] = "two"
	nums[3] = "three"

	fmt.Println(nums[1]) // getting value from map

	fmt.Println(nums) // get all in signle time

	fmt.Println(nums[5]) // getting nothing - its depend on value data type int=0 string ="" bool=false
	if nums[5] == ""{
		fmt.Println("there is nothing")
	}else{	
		fmt.Println("there is something")
	}

	fmt.Println(len(nums))

	delete(nums,1) // delete the single key and value

	clear(nums) // clear the all map data

	nums[1] = "one"
	nums[2] = "two"
	nums[3] = "three"

	// v,b := nums[121] // map return 1st value and 2nd boolean data | data present data retrun otherwise zeroed data
	v,ok := nums[2] // map return 1st value and 2nd boolean data 
	fmt.Println(v)
	fmt.Println(ok)

	killerName1 := map[string]string{"john wick":"babayaga","mushashi":"invincible","guts":"blackswordsman"} // inline map
	killerName2 := map[string]string{"john wick":"babayaga","mushashi":"invincible","guts":"blackswordsman"} // inline map
	fmt.Println(maps.Equal(killerName1,killerName2))



}