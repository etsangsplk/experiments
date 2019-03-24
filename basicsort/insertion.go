package main

import (
	"fmt"
	"sort"

	"./helper"
)

func main() {
	for i := 0; i < 10; i++ {
		a := helper.GenerateSlice(i)
		b := make([]int, len(a))
		copy(b, a) // copy to b from a
		insertionSort(a)
		sort.Ints(b)

		for k, v := range b {
			if v != a[k] {
				fmt.Printf("different at :%d", k)
			}
		}
		fmt.Printf("a: %v \nb: %v\n", a, b)
	}
}

// Maintain a sorted and "unknown" sections.
// Sorted section started at 0
// Try to find the smallest element from the unknown sections
// and add to the "sorted" section.
func insertionSort(a []int) {
	n := len(a)
	i := 0
	for i < n {
		j := i + 1
		for j < n {
			if a[i] > a[j] {
				helper.Swap(&a[i], &a[j])
			}
			j++
		}
		i++
	}

}
