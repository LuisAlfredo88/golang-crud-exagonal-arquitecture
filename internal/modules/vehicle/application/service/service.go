package service

import (
	vehicleRepo "golang-crud-exagonal-arquitecture/internal/modules/vehicle/model/repository"
	vehicleService "golang-crud-exagonal-arquitecture/internal/modules/vehicle/model/service"
)

type vehicle struct {
	vehicleRepository vehicleRepo.VehicleRepository
}

func NewVehicleService(vehicleR vehicleRepo.VehicleRepository) vehicleService.VehicleService {
	return &vehicle{
		vehicleRepository: vehicleR,
	}
}
