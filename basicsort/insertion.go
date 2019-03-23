package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		a := generateSlice(i)
		iSort(a)
		fmt.Printf("a: %v", a)
	}
}

//
func iSort(a []int) {
	n := len(a)
	i := 0
	for i < n {
		j := i + 1
		for j < n {
			if a[i] > a[j] {
				swap(&a[i], &a[j])
				// a[i], a[j] = a[j], a[i]
			}
			j++
		}
		i++
	}

}

//
func swap(a *int, b *int) {
	*a, *b = *b, *a
}

// Generates a slice of size, size filled with random numbers
func generateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}
