package main

import ( "fmt")

type VehiclePrinter interface {
	printMe()
}


type Vehicle struct {
	licensePlate string
	vehicletype string
}

func (v *Vehicle) LicensePlate()  string {
	return fmt.Sprintf("LicensePlate: \n", v.licensePlate)
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
		licensePlate: "1234",
		vehicletype: "car",
		wheels: 4,
		comment: "i am a car",
	}
}

func (c *Car) printMe()  {
	fmt.Printf("licensePlate: %v car details: %v\n", c.LicensePlate(), c.String())
}
	

type Bus struct {
	Vehicle
	FourWheeler
	comment string
}

func NewBus() *Bus{
	return &Bus {
		licensePlate: "5678",
		vehicletype: "bus",
		wheels: 2,
		comment: "i am a bus",
	}
}


func (b *Bus) printMe() {
	fmt.Printf("bus details : %v\n", b)
}