package main

import "fmt"

type T interface {
	PostOrder()
	PreOrder()
	InOrder()
	Size() int
	Diameter() int
	IsCycle() bool
	IsEmpty() bool
	Mirror()
	ToList()
}

type BTree1 struct {
	T
	fmt.Stringer
}

func (t *BTree1) NewBTree1() *BTree1 {
	return BTree1{}
}

func (t *BTree1) ToString() string {
	return
}
