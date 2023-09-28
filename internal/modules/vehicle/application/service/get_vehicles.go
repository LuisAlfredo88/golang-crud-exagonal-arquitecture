package service

import (
	"golang-crud-exagonal-arquitecture/internal/modules/vehicle/model/entity"
)

func (l *vehicle) GetAllVehicles() ([]entity.Vehicle, error) {
	return l.vehicleRepository.GetAllVehicles()
}
