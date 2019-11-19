package main

type Item util.Item

type L interface {
	Add(x Item)
	DeleteAt(x int)
	GetItemAt(x int) Item
}

type List1 struct {
	L
}

func NewList(size int) *List1 {

}
