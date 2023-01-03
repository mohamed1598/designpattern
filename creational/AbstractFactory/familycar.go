package abstractfactory

type FamilyCar struct{}

func (*FamilyCar) NumDoors()int{
	return 5
}
func (*FamilyCar) NumSeats()int{
	return 5
}
func (*FamilyCar) NumWheels()int{
	return 4
}