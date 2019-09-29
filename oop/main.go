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


	//Ceiling operator
	// result already rounded down by integer division
	fmt.Printf("x is odd math ceil is %v \n", math.Ceil(float64(3/2)))
	fmt.Printf("x is even math ceil is %v \n", math.Ceil(float64(4/2)))

	// normal integerter division also floored the result
	fmt.Printf("x is odd nothing is %v\n", (3/2))
	fmt.Printf("x is even nothing is %v \n", (4/2))

	// normal integerter division also floored the result
	// float should not round down 
	fmt.Printf("x is odd nothing is float64 all to get precision  %v\n", float64(3)/float64(2))
	fmt.Printf("x is even nothing is float64 all to get preiceison %v \n", float64(4)/float64(2))

	fmt.Printf("x is odd nothing is float64 all then math Ceil %v\n", math.Ceil(float64(3)/float64(2)))
	fmt.Printf("x is even nothing is float64 all then math Ceil %v \n", math.Ceil(float64(4)/float64(2)))

	fmt.Printf("x is odd nothing is float64 all then math Floor %v\n", math.Floor(float64(3)/float64(2)))
	fmt.Printf("x is even nothing is float64 all then math Floor %v \n", math.Floor(float64(4)/float64(2)))



}
