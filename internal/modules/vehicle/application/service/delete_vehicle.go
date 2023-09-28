package service

import (
	"errors"
)

func (p *vehicle) DeleteVehicle(vehcileId int32) error {
	if vehcileId <= 0 {
		return errors.New("must specify vehcile id")
	}

	logError := p.vehicleRepository.DeleteVehicle(vehcileId)
	return logError
}
