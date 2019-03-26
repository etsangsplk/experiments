package main

import (
	"fmt"
)

func main() {
	a := []int{}
	x := twoSum(a, 0)
	fmt.Printf("Index %v \n", x)
	a = []int{1}
	x = twoSum(a, 1)
	fmt.Printf("Index %v \n", x)
	a = []int{1, 1}
	x = twoSum(a, 2)
	fmt.Printf("Index %v \n", x)

	a = []int{1, 1, 3}
	x = twoSum(a, 4)
	fmt.Printf("Index %v \n", x)

	a = []int{1, -1, 3}
	x = twoSum(a, 2)
	fmt.Printf("Index %v \n", x)
}

func twoSum(a []int, target int) []int {
	n := len(a)
	i := 0

	for i < n {
		j := i + 1
		for j < n {
			if i == j {
				continue
			}
			if a[i]+a[j] == target {
				return []int{i, j}
			}
			j++
		}
		i++
	}
	return []int{}
}
