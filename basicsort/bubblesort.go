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
		bubbleSort(a)
		sort.Ints(b)

		for k, v := range b {
			if v != a[k] {
				fmt.Printf("different at :%d \n", k)
			}
		}
		fmt.Printf("a: %v \nb: %v\n", a, b)
	}
}

/* For each pass, swap the whole array to less wrong configuration by:
 * Partition array into sorted and unsorted,
 * Bubble up largest element to the end of the unsorted array
 * Now shrink the unsorted by one and expand the sorted section at end by one
 * Repeat until the unsorted is of size one or less
 */
func bubbleSort(a []int) {
	n := len(a)
	i := 0
	for i < n-1 {
		// The inner loop is for swapping or bubbling the largest element to end of unsorted array
		// Each time starts from first element, till to end of jsut beofre the srted array.
		j := 0
		for j < n-1-i {
			if a[j] > a[j+1] {
				helper.Swap(&a[j], &a[j+1])
			}
			j++
		}
		//
		i++
	}
}
