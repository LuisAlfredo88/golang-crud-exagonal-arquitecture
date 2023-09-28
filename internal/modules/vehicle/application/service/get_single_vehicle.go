package service

import (
	"golang-crud-exagonal-arquitecture/internal/modules/vehicle/model/entity"
)

func (l *vehicle) GetVehicle(vehicleId int32) (entity.Vehicle, error) {
	return l.vehicleRepository.GetVehicle(vehicleId)
}
