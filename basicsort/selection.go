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
		selectionSort(a)
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
// Sorted section started at 0.
// Invariant S is sorted.
// Find the smallest element from the unknown sections
// and add to the "sorted" section at the end.
// Repeat until there is 1 left in unsorted list.
func selectionSort(a []int) {
	n := len(a)
	i := 0 // i: sorted section
	for i < n {
		j := i + 1 // j: unsorted section
		for j < n {
			if a[i] > a[j] {
				helper.Swap(&a[i], &a[j]) // swap or append to end of sorted section
			}
			j++
		}
		i++
	}

}

func selectionSortWithRecursion(a []int) {

}
