package builder

type BusBuilder struct{
	vehicle VehicleProduct
}

func (b *BusBuilder) SetWheels()BuildProcess{
	b.vehicle.Wheels = 8
	return b
}
func (b *BusBuilder) SetSeats()BuildProcess{
	b.vehicle.Seats = 30
	return b
}
func (b *BusBuilder) SetStructure()BuildProcess{
	b.vehicle.Structure = "Bus"
	return b
}
func (b *BusBuilder) GetVehicle() VehicleProduct{
	return b.vehicle
}