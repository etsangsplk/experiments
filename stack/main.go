package main

import "fmt"

type bin int

func (b bin) String() string {
	return fmt.Sprintf("%b", b)
}

func main() {
	s := New(0)
	fmt.Printf("%v\n", s.Pop())
	fmt.Printf("%v\n", s.Size())
	fmt.Printf("%v\n", s.IsFull())
	fmt.Printf("%v\n", s.IsEmpty())
	s = New(1)
	s.Push(bin(1))
	fmt.Printf("%v\n", s.Pop())
	fmt.Printf("%v\n", s.Size())
	fmt.Printf("%v\n", s.IsFull())
	fmt.Printf("%v\n", s.IsEmpty())

	s2 := NewStack2()
	fmt.Printf("%v\n", s2.Pop())
	fmt.Printf("%v\n", s2.Size())
	fmt.Printf("%v\n", s2.IsEmpty())

	s2 = NewStack2()
	s2.Push(bin(1))
	fmt.Printf("%v\n", s2.Pop())
	fmt.Printf("%v\n", s2.Size())
	fmt.Printf("%v\n", s2.IsEmpty())
}
