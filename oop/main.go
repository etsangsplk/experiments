package main

import (
	"fmt"
	"math"
)

func main() {
	c := NewCar()
	b := NewBus()
	// program against the interface
	vl := []VehiclePrinter{c, b}

	for k, v := range vl {
		fmt.Printf("k :%v v: %v\n ", k, v.printMe())
	}

	fmt.Printf("annonymout struct: %v \n", struct {
		age  int
		name string
	}{age: 10, // have to repeat definition of anonymout struct if more than 1 fields
		name: "wawa"})

	fmt.Println(func() string { return "ddmfkdmfdmg" }())
	fmt.Println(func(x string) string { return fmt.Sprintf("yyyyy%s", x) }("ddmfkdmfdmg"))

	//Floor operator
	fmt.Printf("x is odd math floor is %v  \n", math.Floor(float64(3/2)))

	fmt.Printf("x is even math floor is %v \n", math.Floor(float64(4/2)))

}
