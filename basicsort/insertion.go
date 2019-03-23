package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		a := generateSlice(i)
		b := make([]int, len(a))
		copy(b, a) // copy to b from a
		iSort(a)
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
				swap(&a[i], &a[j])
			}
			j++
		}
		i++
	}

}

// helper swap function.
func swap(a *int, b *int) {
	*a, *b = *b, *a
}

// generates a slice of size, size filled with random numbers
func generateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}
