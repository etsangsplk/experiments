package main

import ( "fmt")

type VehiclePrinter interface {
	printMe() string
}


type Vehicle struct {
	LicensePlate string
	Vehicletype string
}

func (v *Vehicle) PLicensePlate() string {
	return fmt.Sprintf("LicensePlate: %v\n", v.LicensePlate)
}

type FourWheeler struct {
	wheels int
}

func (f *FourWheeler) String() string {
	//f.wheels or else inifite loop
	return fmt.Sprintf("FourWheeler = %v\n", f.wheels)
}

type Car struct {
	Vehicle
	FourWheeler
	comment string
}

func NewCar() *Car{
	return &Car{
		// embedded annoymous struct available via type only
		Vehicle: Vehicle{ LicensePlate: "1234",
						  Vehicletype: "car",},
		FourWheeler: FourWheeler{ wheels: 4,},
		comment: "i am a car",
	}
}

func (c *Car) printMe() string {
	return fmt.Sprintf("licensePlate: %v car details: %v\n", c.PLicensePlate(), c.String())
}
	

type Bus struct {
	Vehicle
	FourWheeler
	comment string
}

func NewBus() *Bus{
	return &Bus {
		Vehicle:  Vehicle{ LicensePlate: "5678",
		 		   Vehicletype: "bus",},
		FourWheeler:  FourWheeler{wheels: 2,},
		comment: "i am a bus",
	}
}


func (b *Bus) printMe() string {
	return fmt.Sprintf("bus details : %v\n", b)
}