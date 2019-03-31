package main

import (
	"fmt"
	"sync"
)

var coeff []int
var x int

func main() {
	once := &sync.Once{}
	var n int
	fmt.Println("n:10 x:1 ans should be 56")
	fmt.Scanf("n:%d x:%d", &n, &x)
	once.Do(func() { fillCofficient(n) })
	fmt.Printf("coefficient: %v\n", coeff)
	ret := polynomial(n)
	fmt.Printf("result: %v iteration: %d\n", ret, n)

	fmt.Printf("result: %v iteration: %d\n", polynomialArray(n), n)
}

// Brutal force
// Here
func polynomial(n int) int {
	if n == 0 {
		return coeff[0]
	}
	return coeff[n]*x + polynomial(n-1)
}

func polynomialArray(n int) int {
	var ret int
	for i := 0; i <= n; i++ {
		ret = ret + coeff[i]*x
	}
	return ret
}

func fillCofficient(n int) {
	coeff = make([]int, n+1)
	for i := 0; i <= n; i++ {
		coeff[i] = i
	}

	coeff[0] = 1
}
