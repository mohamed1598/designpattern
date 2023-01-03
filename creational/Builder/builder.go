package builder

type BuildProcess interface{
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	GetVehicle() VehicleProduct
}

type VehicleProduct struct{
	Wheels int
	Seats int
	Structure string
}

type ManufacturingDirector struct{
	builder BuildProcess
}

func (f *ManufacturingDirector) Construct(){
	f.builder.SetSeats().SetStructure().SetWheels()
}

func (f *ManufacturingDirector) SetBuilder(b BuildProcess){
	f.builder = b
}