package main

import (
	"fmt"
	"strings"

	"../basicsort/helper"
)

func main() {

	d := []string{}

	for i := 0; i < 20; i++ {
		d = append(d, helper.GenerateString(i))
	}

	for _, s := range d {
		fmt.Printf("string %s\n", s)
		l := reverse(s)
		r := ReverseRecur(s)
		if !strings.EqualFold(l, r) {
			fmt.Printf("reverse string %s recurrsive\n", l, r)
		} else {
			fmt.Printf("pass: %s\n", s)
		}
	}
}

// Note:
// * Golang String in golang is immutatble
// * Need to manipulate the underlying Rune
func reverse(s string) string {
	n := len(s)
	i := 0
	j := n - 1
	temp := []rune(s)
	//temp := make([]rune, n)

	for i <= j {
		temp[i], temp[j] = rune(s[j]), rune(s[i])
		i++
		j--
	}
	return string(temp)
}

func ReverseRecur(s string) string {
	r := []rune(s)
	i := 0
	j := len(r) - 1
	reverseRecur(r, i, j)
	return string(r)
}

func reverseRecur(r []rune, i int, j int) {
	if i < j {
		r[i], r[j] = r[j], r[i]
		i++
		j--
		reverseRecur(r, i, j)
	}
}
