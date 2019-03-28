package main

import (
	"fmt"
)

func main() {
	s := "abcde"
	fmt.Printf("string %s reverse string %s\n", s, reverse(s))

	s = ""
	fmt.Printf("string %s reverse string %s\n", s, reverse(s))

	s = "z"
	fmt.Printf("string %s reverse string %s\n", s, reverse(s))

}

// Note:
// * Golang String in golang is immutatble
// * Need to manipulate the underlying Rune
func reverse(s string) string {
	n := len(s)
	i := 0
	j := n - 1
	temp := make([]rune, n)

	for i <= j {
		temp[i] = rune(s[j])
		temp[j] = rune(s[i])
		i++
		j--
	}
	return string(temp)
}

//func reverseRecur(s string, n int) string {

//}
