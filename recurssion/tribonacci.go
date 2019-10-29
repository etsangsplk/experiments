package main

import "fmt"

var result = make([]int, 37)

func init() {
	result[0] = 0
	result[1] = 1
	result[2] = 1
}

func tribonacci(n int) int {
	if n > 3 {
		var n3, n2, n1 int
		if 0 != result[n-3] {
			n3 = result[n-3]
		} else {
			n3 = tribonacci(n - 3)
			result[n-3] = n3
		}
		if 0 != result[n-2] {
			n2 = result[n-2]
		} else {
			n2 = tribonacci(n - 2)
			result[n-2] = n2
		}
		if 0 != result[n-1] {
			n1 = result[n-1]
		} else {
			n1 = tribonacci(n - 1)
			result[n-1] = n1
		}
	}
	return result[n-3] + result[n-2] + result[n-1]
}

func main() {
	n := 25
	fmt.Println(tribonacci(n))
	fmt.Println(tribonacci(0))
	fmt.Println(tribonacci(2))
	fmt.Println(tribonacci(36))
}
