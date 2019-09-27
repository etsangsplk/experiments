package main

import ( "fmt")

type VehiclePrinter interface {
	printMe()
}


type Vehicle struct {
	LicensePlate string
	Vehicletype string
}

func (v *Vehicle) PLicensePlate()  string {
	return fmt.Sprintf("LicensePlate: \n", v.LicensePlate)
}

//func (v *Vehicle) String() {
//	fmt.Printf("Vehicle: \n", v)
//}


type FourWheeler struct {
	wheels int
}

func (f *FourWheeler) String() string {
	return fmt.Sprintf("FourWheeler = %v\n", f)
}

type Car struct {
	Vehicle
	FourWheeler
	comment string
}

func NewCar() *Car{
	return &Car{
		Vehicle.LicensePlate: "1234",
		Vehicle.Vehicletype: "car",
		wheels: 4,
		comment: "i am a car",
	}
}

func (c *Car) printMe()  {
	fmt.Printf("licensePlate: %v car details: %v\n", c.PLicensePlate(), c.String())
}
	

type Bus struct {
	Vehicle
	FourWheeler
	comment string
}

func NewBus() *Bus{
	return &Bus {
		Vehicle.LicensePlate: "5678",
		Vehicle.Vehicletype: "bus",
		wheels: 2,
		comment: "i am a bus",
	}
}


func (b *Bus) printMe() {
	fmt.Printf("bus details : %v\n", b)
}