package builder

import (
	"fmt"
	"testing"
)

func TestBuilderPattern(t *testing.T){
	 TestCarBuilder(t)
	 TestBikeBuilder(t)
	 TestBusBuilder(t)
}
func TestCarBuilder(t *testing.T){
	fmt.Println("Test1 : test car builder")
	manufacturingComplex := &ManufacturingDirector{}
	carBuilder := &CarBuilder{}
	manufacturingComplex.SetBuilder(carBuilder)
	manufacturingComplex.Construct()
	car := carBuilder.GetVehicle()
	if car.Wheels != 4{
		t.Errorf("Wheels on a car must be 4 and they were %d\n",car.Wheels)
	}
	if car.Structure != "Car"{
		t.Errorf("Structure on a car must be 'Car' and was %s\n",car.Structure)
	}
	if car.Seats != 5{
		t.Errorf("Seats on a car must be 5 and they were %d\n",car.Seats)
	}
}
func TestBikeBuilder(t *testing.T){
	fmt.Println("Test2 : test motorbike car builder")
	manufacturingComplex := &ManufacturingDirector{}
	BikeBuilder := &BikeBuilder{}
	manufacturingComplex.SetBuilder(BikeBuilder)
	manufacturingComplex.Construct()
	motorbike := BikeBuilder.GetVehicle()
	if motorbike.Wheels != 2{
		t.Errorf("Wheels on a motorbike must be 2 and they were %d\n",motorbike.Wheels)
	}
	if motorbike.Structure != "Motorbike"{
		t.Errorf("Structure on a motorbike must be 'Motorbike' and was %s\n",motorbike.Structure)
	}
	if motorbike.Seats != 1{
		t.Errorf("Seats on a motorbike must be 1 and they were %d\n",motorbike.Seats)
	}
}
func TestBusBuilder(t *testing.T){
	fmt.Println("Test3: test bus builder")
	manufacturingComplex := &ManufacturingDirector{}
	busBuilder := &BusBuilder{}
	manufacturingComplex.SetBuilder(busBuilder)
	manufacturingComplex.Construct()
	bus := busBuilder.GetVehicle()
	if bus.Wheels != 8{
		t.Errorf("Wheels on a bus must be 8 and they were %d\n",bus.Wheels)
	}
	if bus.Structure != "Bus"{
		t.Errorf("Structure on a bus must be 'Bus' and was %s\n",bus.Structure)
	}
	if bus.Seats != 30{
		t.Errorf("Seats on a bus must be 30 and they were %d\n",bus.Seats)
	}
}