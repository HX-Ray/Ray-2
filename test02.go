package main

import(
	"fmt"
)

func  RomanToInt(romanNum string) int {
	var sum = 0
	var roman = map[byte]int{'I':1 , 'V':5 , 'X':10 , 'L':50 , 'C':100 , 'D':500 , 'M':1000}

	for i:=0; i<len(romanNum)-1; i++ {
		word := romanNum[i]
		wordnext := romanNum[i+1]

		if  roman[word] < roman[wordnext] && roman[wordnext] <= roman[word]*10 {
			sum += roman[word]*(-1)
		} else {
			sum += roman[word]
		}
	}
	sum += roman[romanNum[len(romanNum)-1]]

	return sum
}


func main(){
	var alb int
	var romanNum string
	fmt.Scanf("%s", &romanNum)

	alb = RomanToInt(romanNum)
	fmt.Println(alb)
}
