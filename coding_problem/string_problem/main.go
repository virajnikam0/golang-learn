package main

import (
	"fmt"
	// "internal/runtime/maps"
	"strconv"
	"strings"
)

// world chess champinship
func ChessChampionsship(gamesStr string, points int) int {

	/*
		tatol game = 14
		win(2) loose(0) draw(1)
		c = person1
		n = person2
		d = draw matches

		if win > loose {
			win * 2 * 60
			loose * 2 * 40
		}
		if win == loose {
			c * 2 * 55
			n * 2 * 45
		}



	*/

	var p1, p2, drawCount int = 0, 0, 0

	for _, singleGame := range gamesStr {
		if singleGame == 'C' {
			p1 = p1 + 1
		} else if singleGame == 'N' {
			p2 = p2 + 1
		} else {
			drawCount = drawCount + 1
		}
	}

	person1Won := p1*2 + drawCount*1

	person2Won := p2*2 + drawCount*1

	if person1Won > person2Won {
		return 60 * points
	}
	if person1Won == person2Won {
		return 55 * points
	}

	return 40 * points

}

// date checking format
func DateChecking(param string) string {

	if len(param) == 10 {
		// get char 1 2 and 4 5
		// if char is > 12 use logic
		var tempDate1 string = ""
		var tempDate2 string = ""

		for index, _ := range param {
			switch index {
			case 0, 1:
				tempDate1 = tempDate1 + string(param[index])
			case 3, 4:
				tempDate2 = tempDate2 + string(param[index])
			}
		}

		date1, _ := strconv.Atoi(tempDate1)
		date2, _ := strconv.Atoi(tempDate2)

		if date1 <= 12 && date2 <= 12 {
			return "Both"
		} else if date1 <= 12 && date2 > 12 {
			return "MM/DD/YYYY"
		} else {
			return "DD/MM/YYYY"
		}

	}

	return " "

}

// find largest num from string
func LargeAddSubstring(param string) int {
	var largestOddNum int = 0

	for index, data := range param {
		if (data-'0')%2 == 1 {
			temp, _ := strconv.Atoi(param[:index+1])

			if largestOddNum < temp {
				largestOddNum = temp
			}
		}

	}
	return largestOddNum
}

func ConvertFistCharUpper(param string) string {

	arrString := strings.Split(param, " ")

	for index, data := range arrString {
		if strings.ToUpper(data) == data {
			continue
		} else {
			tempVar := strings.ToLower(data)

			tempVar = strings.ToUpper(tempVar[:1]) + strings.ToLower(tempVar[1:])

			arrString[index] = tempVar

		}
	}

	newStr := strings.Join(arrString, " ")

	fmt.Println(newStr)

	return newStr

}

/*
check jeff can read the lette hat present iside the word if yes then write yes if no he dont then no
*/
func JeffCanReadLetter(param1 int, param2 []string, param3 string) []string {

	letterToKnow := make(map[string]rune)
	// ceating yes/no array
	arrYesNo := []string{}

	for _, data := range param3 {
		_, ok := letterToKnow[string(data)]
		if ok {
			continue
		} else {
			letterToKnow[string(data)] = data

		}
	}

	for i := range param1 {
		currLetter := param2[i]
		temp := 0
		tempNo := len(currLetter)
		for _, data := range currLetter {
			_, ok := letterToKnow[string(data)]
			if ok {
				temp++
				continue
			} else {
				break
			}
		}
		if tempNo == temp {
			arrYesNo = append(arrYesNo, "Yes")
		} else {
			arrYesNo = append(arrYesNo, "No")
		}
	}

	return arrYesNo

}

// chnage the ballon color
func ChnageBalonColor(param1 string) int {

	var c1, c2 int = 0, 0

	for i := 0; i <= len(param1)/2; i++ {
		if param1[i] == 'a' {
			c1++
		} else {
			c2++
		}
		if param1[len(param1)-1-i] == 'b' {
			c2++
		} else {
			c1++
		}
	}

	if c1 > c2 {
		return c2
	}
	return c1
}

// does light blub is on or off
func LightBlubOnOrOff(param1 string, param2 string) bool {
	count := 0
	for i := 0; i <= len(param1)/2; i++ {
		if param1[i] != param2[i] || param1[len(param1)-1-i] != param2[len(param2)-1-i] {
			count++
		}
	}

	if count%2 == 1 {
		return false
	} else {
		return true
	}
}

// girls play fairly or not
func GirlsPlayPianofairlyOrNot(param1 string) bool {
	count := 0
	for i := 0; i < len(param1)/2; i += 2 {

		if len(param1) > 2 {
			if param1[i] == 'A' {
				if param1[i+1] == 'B' {
					count++
				}
			}

			if param1[i] == 'B' {
				if param1[i+1] == 'A' {
					count++
				}
			}

			if param1[len(param1)-1-i] == 'A' {
				if param1[len(param1)-2-i] == 'B' {
					count++
				}
			}

			if param1[len(param1)-1-i] == 'B' {
				if param1[len(param1)-2-i] == 'A' {
					count++
				}
			}
		} else {
			if param1[i] == 'A' {
				if param1[i+1] == 'B' {
					count++
				}
			}

			if param1[i] == 'b' {
				if param1[i+1] == 'A' {
					count++
				}
			}
		}

	}

	if (len(param1) / 2) == count {
		return true
	} else {
		return false
	}
}

func main() {
	// fmt.Println(ChessChampionsship("CDCDCDCDCDCDCD", 400))

	// fmt.Println(DateChecking("21/05/2001"))
	// fmt.Println(DateChecking("10/15/2069"))
	// fmt.Println(DateChecking("10/18/2069"))
	// fmt.Println(DateChecking("05/11/1999"))
	// fmt.Println(DateChecking("29/02/2024"))

	// fmt.Println(LargeAddSubstring("987654321"))
	// fmt.Println(LargeAddSubstring("233564"))

	// fmt.Println(ConvertFistCharUpper("hello world"))
	// fmt.Println(ConvertFistCharUpper("this is a CODECHEF problem"))
	// fmt.Println(ConvertFistCharUpper("this is a CODECHEF N.A."))

	// fmt.Println(JeffCanReadLetter(3, []string{"abc", "bbc", "xxx"}, "abc"))

	// fmt.Println(ChnageBalonColor("abbaaab"))

	// fmt.Println(LightBlubOnOrOff("00", "11"))

	// fmt.Println(GirlsPlayPianofairlyOrNot("AA"))

	fmt.Println(strings.Contains("11111110", "010"))
	fmt.Println(strings.Contains("11111110", "101"))

}
