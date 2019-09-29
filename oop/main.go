package main

import ("fmt")

func main() {
	c := NewCar()
	b := NewBus()
	// program against the interface 
	vl := []VehiclePrinter{c, b}

	for k, v := range vl {
		fmt.Printf("k :%v v: %v\n ", k, v.printMe())
	}


	fmt.Printf("annonymout struct: %v \n", struct {
		   age int
		   name string
		}{ age: 10,  // have to repeat definition of anonymout struct if more than 1 fields
		   name: "wawa",})
}