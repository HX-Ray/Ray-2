package main

import(
	"fmt"
)

func  Replace(s string, target, replace byte) string {
	var slice []byte

	for i:=0; i<len(s); i++ {

		if s [i] == target {
			slice = append(slice, replace)
		} else {
			slice = append(slice, s[i])
		}
	}

	return string(slice)
}

func main(){
	var s string
	var target, replace byte

	fmt.Scanf("%s", &s)
	fmt.Scanf("%c\n", &target)
	fmt.Scanf("%c\n", &replace)

	s = Replace(s, target, replace)
	fmt.Printf("%s\n", s)
}

