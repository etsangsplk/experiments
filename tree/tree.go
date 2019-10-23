package main

type Tree interface {
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
