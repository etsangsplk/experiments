package main

import (
	"fmt"
)

type bin int

func (b bin) String() string {
	return fmt.Sprintf("%b", b)
}
func main() {
	q := New(10)
	fmt.Printf("q: %v\n ", q)

	q2 := NewQ2()
	for i := 0; i < 10; i++ {
		q2.Enqueue(bin(i))
	}
	fmt.Printf("q2: %d\n ", q2.Items)
}
