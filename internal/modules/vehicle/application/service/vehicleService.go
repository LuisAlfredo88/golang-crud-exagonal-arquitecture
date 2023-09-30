package service

import (
	"errors"
	"golang-crud-exagonal-arquitecture/internal/modules/vehicle/model/entity"
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

func (p *vehicle) DeleteVehicle(vehcileId int32) error {
	if vehcileId <= 0 {
		return errors.New("must specify vehcile id")
	}

	logError := p.vehicleRepository.DeleteVehicle(vehcileId)
	return logError
}

func (l *vehicle) GetVehicle(vehicleId int32) (entity.Vehicle, error) {
	return l.vehicleRepository.GetVehicle(vehicleId)
}

func (l *vehicle) GetAllVehicles() ([]entity.Vehicle, error) {
	return l.vehicleRepository.GetAllVehicles()
}

func (p *vehicle) SaveVehicle(vehicleEntity entity.Vehicle) error {
	err := error(nil)

	if vehicleEntity.ID > 0 {
		err = entity.ValidateVehicle(vehicleEntity)
	} else {
		err = entity.ValidateVehicle(vehicleEntity)
	}

	if err != nil {
		return err
	}

	logError := p.vehicleRepository.SaveVehicle(vehicleEntity)
	return logError
}
