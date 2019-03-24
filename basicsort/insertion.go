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
				fmt.Printf("different at :%d \n", k)
			}
		}
		fmt.Printf("a: %v \nb: %v\n", a, b)
	}
}

// Partition list into sorted and unsorted sections.
// Sorted section starts as 1 single element.
// Empty list is ignored
// Invariant is S is sorted.
// Just grab an element from Unsorted list
// * comparing this element to sorted list starting from last element of sorted list
// * swap the previous element with current postion to make space
// * until the postion where the element infront is smaller
// Repeat until unsorted list is of 1 element
func insertionSort(a []int) {
	n := len(a)

	// start at 1, empty list of size 0.
	// for-loop will not be executed.
	i := 1
	for i < n {
		j := i
		// Expand the sorted list boundary by 1, including 1 element of unsorted list
		// Starting at end of sorted list.
		for j > 0 {
			// Compare myself against previous element
			if a[j] < a[j-1] {
				helper.Swap(&a[j], &a[j-1])
			}
			j-- //move point to element
		}
		i++
	}

}
