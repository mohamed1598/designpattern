package abstractfactory

import "fmt"

type Motorbike interface {
	GetMotorbikeType() int
}

const (
	SportMotorbikeType  = 1
	CruiseMotorbikeType = 2
)

type MotorbikeFactory struct{}

func (*MotorbikeFactory) Build(v int) (Vehicle, error) {
	switch v {
	case SportMotorbikeType:
		return new(SportMotorbike), nil
	case CruiseMotorbikeType:
		return new(CruiseMotorbike), nil
	default:
		return nil, fmt.Errorf("vehicle of type %d not recognized",v)
	}
}