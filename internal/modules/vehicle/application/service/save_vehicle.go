package service

import (
	"golang-crud-exagonal-arquitecture/internal/modules/vehicle/model/entity"
)

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
