package builder

type CarBuilder struct{
	vehicle VehicleProduct
}

func (c *CarBuilder) SetWheels() BuildProcess{
	c.vehicle.Wheels =4
	return c
}
func (c *CarBuilder) SetSeats() BuildProcess{
	c.vehicle.Seats = 5
	return c
}
func (c *CarBuilder) SetStructure() BuildProcess{
	c.vehicle.Structure = "Car"
	return c
}
func (c *CarBuilder) GetVehicle() VehicleProduct{
	return c.vehicle
}