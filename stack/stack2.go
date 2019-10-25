package main

import "github.com/etsangsplk/experiments/util"

type Item2 util.Item

type Stack2 interface {
	Pop() Item2
	Push(i Item2)
	Size() int
}

type S2 struct {
	Stack2
	items []Item2
}

func NewStack2() *S2 {
	return &S2{
		items: make([]Item2, 0),
	}
}

func (s *S2) Pop() Item2 {
	if s.IsEmpty() {
		return nil
	}
	endIndex := len(s.items) - 1
	x := s.items[endIndex]
	s.items = s.items[:endIndex]
	return x
}

func (s *S2) Push(i Item2) {
	s.items = append(s.items, i)
}

func (s *S2) Size() int {
	return len(s.items)
}

func (s *S2) IsEmpty() bool {
	return q.Size() == 0
}
