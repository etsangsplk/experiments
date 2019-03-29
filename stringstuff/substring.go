package main

import (
	"fmt"

	"../basicsort/helper"
)

func main() {
	data := []string{}
	n := 10
	for i := 0; i < n; i++ {
		data = append(data, helper.GenerateString(i))
	}

	for j := 0; j < n; j++ {
		fmt.Printf("%s\n", data[j])

	}

	//p := []rune(string("abcde"))
	//q := []rune(string("abcde"))

}

// Returns an array of tuples of where the substrings occurrs in source string.
// If none is found returns empty tuple.
func substring(source string, sub string) []int {
	n := len(source)
	m := len(sub)
	ret := []int{}
	if m > n {
		return []int{}
	}

	if m == 0 {
		return ret
	}

	i := n
	j := m
	for i > 0 {
		for j > 0 {
			if sub[j] == source[i] {
				j--
				i--
			}
		}
	}

	return ret
}

// Returns an array of strings tokenized by seperators.
// If none is found, returns empty tuple
func splitString(s string, sep string) []string {
	n := len(s)
	//m := len(sep)
	ret := []string{}
	i, j, k := 0, 0, 0
	for i < n {
		j := i
		for j < n {
			//f

		}

	}
	return ret
}

// Replace all occurrences of with string from s string with X.
func replaceString(s string, tomask string) string {
	n := len(s)
	m := len(tomask)
	if m > n {
		return ""
	}

	for i:=0; i< n-m-1; i++ {
		for j:=0; j< m-1; j++ {
			if s[i] != tomask[j] {
				break; // skip the comparison agains tomask, 
				// do next comparison on i++
			}
		}
	}


	return ""
}

// a convenience functions for testing
func equals(l, r string) bool {
	i := len(l)
	j := len(r)
	if i != j {
		return false
	}
	for i > 0 {
		if l[i] != r[j] {
			return false
		}
		i--
		j--
	}
	return true
}
