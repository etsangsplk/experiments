package main

import "fmt"

func tribonacci(n int) int {
	switch n {
	case 0:
		return 0
	case 1, 2:
		return 1
	default:
		return tribonacci(n-3) + tribonacci(n-2) + tribonacci(n-1)
	}
}

func main() {
	n := 25
	fmt.Println(tribonacci(n))
	fmt.Println(tribonacci(0))
	fmt.Println(tribonacci(2))
}
