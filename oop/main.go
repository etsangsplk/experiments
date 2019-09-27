package main

import ("fmt")

func main() {
	c := &Car{vehicletype: "car", licensePlate: "1234", comment: "car!!!", wheels: 4}
	b := &Bus{vehicletype: "bus", licensePlate: "5678", comment: "bus!!!"}

	// program against the interface 
	vl := []VehiclePrinter{c, b}

	for k, v := range vl {
		fmt.Printf("k :%v v: %v\n ", k, v)
	}

}