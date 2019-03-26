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
