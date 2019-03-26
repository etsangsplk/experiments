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
		mergeSort(a)
		sort.Ints(b)

		for k, v := range b {
			if v != a[k] {
				fmt.Printf("different at :%d \n", k)
			}
		}
		fmt.Printf("a: %v \nb: %v\n", a, b)
	}
}

func mergeSort(a []int) {

}

func mergeSortR(a []int, i int, j int) {

}

// Merge 2 sorted arrays l and r into a sorted array
func mergeSortedArrays(l []int, r []int) []int {
	n = len(l)
	m = len(r)

	ret = make([]int, n+m)

	i, j, k := n, m, 0

	for j > 0 && i > 0 {
		if l[i] >= r[j] {
			ret[k] = l[i]
			i--
		} else {
			ret[k] = r[j]
			j--
		}
	}

	if i > 0 {
		for i > 0 {
			ret[k] = l[i]
			i--
			k++
		}
	}

	if j > 0 {
		for j > 0 {
			ret[k] = r[j]
			j--
			k++
		}
	}

}
