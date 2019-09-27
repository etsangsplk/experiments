package main

import ("fmt")

func main() {
	c := NewCar()
	b := NewBus()
	// program against the interface 
	vl := []VehiclePrinter{c, b}

	for k, v := range vl {
		fmt.Printf("k :%v v: %v\n ", k, v)
	}

}