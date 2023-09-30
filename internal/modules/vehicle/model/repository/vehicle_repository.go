package repository

import (
	"golang-crud-exagonal-arquitecture/internal/modules/vehicle/model/entity"
)

type VehicleRepository interface {
	SaveVehicle(vehicle entity.Vehicle) error
	GetAllVehicles() ([]entity.Vehicle, error)
	GetVehicle(vehicleId int32) (entity.Vehicle, error)
	DeleteVehicle(vehicle int32) error
}
