package main

import (
	"fmt"
	"strconv"
)

// 2 sum
func twoSum(arr []int, target int) [2]int {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == target {
				return [2]int{i, j}
			}
		}
	}
	return [2]int{-1, -1}
}

// pallindrome num
func pallindromeNum(num int) bool {
	var strNum = strconv.Itoa(num)

	for i := 0; i <= len(strNum)/2; i++ {
		v1 := strNum[i]
		v2 := strNum[len(strNum)-1-i]

		if v1 != v2 {
			return false
		}
	}

	return true
}

// roman number convertor
func romanConv(numString string) int {

	// ceate map
	var romanSymbolValue = make(map[byte]int)
	romanSymbolValue['I'] = 1
	romanSymbolValue['V'] = 5
	romanSymbolValue['X'] = 10
	romanSymbolValue['L'] = 50
	romanSymbolValue['C'] = 100
	romanSymbolValue['D'] = 500
	romanSymbolValue['M'] = 1000

	var globalNum int = 0

	for i := 0; i < len(numString)-1; i++ {
		// if i == len(numString)-1 {
		// 	return 0
		// }

		v1 := numString[i]
		v2 := numString[i+1]

		if v1 == 'I' && v2 == 'V' || v1 == 'I' && v2 == 'X' {
			if v2 == 'V' {
				globalNum += 4
			} else {
				globalNum += 9
			}
		} else if v1 == 'X' && v2 == 'L' || v1 == 'X' && v2 == 'C' {
			if v2 == 'L' {
				globalNum += 40
			} else {
				globalNum += 90
			}
		} else if v1 == 'C' && v2 == 'D' || v1 == 'C' && v2 == 'M' {
			if v2 == 'D' {
				globalNum += 400
			} else {
				globalNum += 900
			}
		} else {

			tempNum := romanSymbolValue[numString[i]]
			globalNum += tempNum
		}

		if i == len(numString)-1 {
			xx := romanSymbolValue[v2]
			globalNum += xx

		}

	}
	return globalNum
}

func main() {

	// 1
	// rtData1 := twoSum([]int{1, 2, 3, 4, 5, 6, 7, 8}, 22)
	// fmt.Println(rtData1)

	// 2
	// rtData2 := pallindromeNum(1231)
	// fmt.Println(rtData2)

	// 3
	rtData3 := romanConv("III")
	fmt.Println(rtData3)

}
