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
}

func substring(s, sub string) bool {

	return false
}

func splitString(s string, sep string) []string {
	return []string{}
}

func replaceString(s string, with string) string {
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
