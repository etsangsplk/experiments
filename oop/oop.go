package main

import ( "fmt")

type VehiclePrinter interface {
	printMe()
}


type Vehicle struct {
	licensePlate string
	vehicletype string
}

func (v *Vehicle) LicensePlate() {
	fmt.Printf("LicensePlate: \n", v.licensePlate)
}

func (v *Vehicle) String() {
	fmt.Printf("Vehicle: \n", v)
}


type FourWheeler struct {
	wheels int
}

func (f *FourWheeler) String() {
	fmt.Printf("FourWheeler = %v\n", f)
}

type Car struct {
	Vehicle
	FourWheeler
	comment string
}


func (c *Car) printMe()  {
	fmt.Printf("licensePlate: %v car details: %v\n", c.LicensePlate(), c.String())
}
	

type Bus struct {
	Vehicle
	comment string
}

func (b *Bus) printMe() {
	fmt.Printf("bu details : %v\n", b.String())
}