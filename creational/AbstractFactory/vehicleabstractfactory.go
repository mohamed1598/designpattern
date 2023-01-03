package abstractfactory

import (
	"fmt"
)

type Vehicle interface {
	NumWheels() int
	NumSeats() int
}

type VehicleFactory interface {
	Build(v int) (Vehicle, error)
}

const (
	CarFactoryType      = 1
	MotorbikeFatoryType = 2
)

func BuildFactory(f int) (VehicleFactory, error) {
	switch f {
	case CarFactoryType:
		return new(CarFactory),nil
	case MotorbikeFatoryType:
		return new(MotorbikeFactory),nil
	default:
		return nil, fmt.Errorf("factory with id %d not recognized", f)
	}
}
